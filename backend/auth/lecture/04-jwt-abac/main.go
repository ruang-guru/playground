// Jalankan kode ini di dalam directory yang sama dengan lokasi kode
// Lalu jalankan semua file .go yang ada di dalamnya dengan command: go run .
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Authorization menggunakan claim JWT
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// JSON body diconvert menjadi creditial struct
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			// return bad request ketika terjadi kesalahan decoding json
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Ambil password dari username yang dipakai untuk login
		expectedPassword, ok := users[creds.Username]

		// return unauthorized jika password salah
		if !ok || expectedPassword.Password != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("user credential invalid"))
			return
		}

		// Deklarasi expiry time untuk token jwt
		expirationTime := time.Now().Add(5 * time.Minute)

		// Buat claim menggunakan variable yang sudah didefinisikan diatas
		claims := &Claims{
			Username: creds.Username,
			Role:     expectedPassword.Role,
			StandardClaims: jwt.StandardClaims{
				// expiry time menggunakan time millisecond
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// return internal error ketika ada kesalahan ketika pembuatan JWT string
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set token string kedalam cookie response
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		w.Write([]byte("Login Success"))
	})

	route.Handle("/profile/{userID:[0-9]+}", AuthMiddleWare("all", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := users[r.Context().Value("props").(*Claims).Username]

		fmt.Println(user)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		params := mux.Vars(r)
		userID, err := strconv.Atoi(params["userID"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if user.ID != userID {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		resp, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	})))

	return route
}

// Test menggunakan curl:
//$ curl -i -X GET http://localhost:8000/profile/1

//$ curl -i -X GET http://localhost:8000/profile/2

// Kita bisa lihat untuk mengakses semua profile yang ada pada data Users
// Kita harus melakukan login terlebih dahulu sesuai dengan user ID yang kita inginkan

// Kita melakukan test login dengan extensi chrome yaitu Talend API Tester: https://chrome.google.com/webstore/detail/talend-api-tester-free-ed/aejoelaoggembcahagimdiliamlcdmfm
