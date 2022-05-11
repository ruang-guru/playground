package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func AuthMiddleWare(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari cookie yang dikirim ketika request
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// return unauthorized ketika token kosong
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Ambil value dari cookie token
		tknStr := c.Value

		// Deklarasi variable claim
		claims := &Claims{}

		//parse JWT token ke dalam claim
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// validasi role
		if role == "all" || claims.Role == role {
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		w.WriteHeader(http.StatusUnauthorized)

	})
}
