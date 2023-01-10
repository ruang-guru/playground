package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// sign dan welcome menggunakan JWT kedalam cookie
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// JSON body diconvert menjadi creditial struct & return bad request ketika terjadi kesalahan decoding json

		// TODO: answer here
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Task: Ambil password dari username yang dipakai untuk login & return unauthorized jika password salah

		// TODO: answer here
		correctPassword := users[creds.Username]
		if correctPassword != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Task: 1. Deklarasi expiry time untuk token jwt
		// 		 2. Buat claim menggunakan variable yang sudah didefinisikan diatas
		//       3. Expiry time menggunakan time millisecond

		// TODO: answer here
		expTime := time.Now().Add(1 * time.Hour)
		claim := &Claims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTime.UnixMilli(),
			},
		}

		// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai

		// TODO: answer here
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

		// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		//       2. Buat return internal error ketika ada kesalahan ketika pembuatan JWT string

		// TODO: answer here
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Task: Set token string kedalam cookie response

		// TODO: answer here
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expTime,
		})
	})

	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		// TODO: answer here
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Task: Ambil value dari cookie token

		// TODO: answer here
		tokenString := cookie.Value

		// Task: Deklarasi variable claim

		// TODO: answer here
		claim := &Claims{}

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada

		// TODO: answer here
		token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
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

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Task: return data dalam claim, seperti username yang telah didefinisikan

		// TODO: answer here
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claim.Username)))
	})

	return mux
}
