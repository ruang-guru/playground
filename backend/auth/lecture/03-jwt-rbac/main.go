// Jalankan kode ini di dalam directory yang sama dengan lokasi kode
// Lalu jalankan semua file .go yang ada di dalamnya dengan command: go run .
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Authorization menggunakan claim JWT
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Public Page"))
	})

	mux.Handle("/admin", AuthMiddleWare("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin Page"))
	})))

	mux.Handle("/student", AuthMiddleWare("student", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Student Page"))
	})))

	mux.Handle("/school", AuthMiddleWare("all", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("School Page"))
	})))

	return mux
}

// Test menggunakan curl:
//$ curl -i -X GET http://localhost:8000/public

//$ curl -i -X GET http://localhost:8000/school

//$ curl -i -X GET http://localhost:8000/admin

//$ curl -i -X GET http://localhost:8000/student

// Kita bisa perhatikan bahwa semua request yang tidak diberikan otorisasi akan mengembalikan status unauthorized dengan code 401
// Karena setiap request akan di cek oleh AuthMiddleWare apaka terdapat token yang dibutuhkan untuk mengakses halaman admin, student dan school menggunakan cookie

// Maka dari itu kita bisa mengakses halaman yang dihandle oleh Middleware dengan melakukan login sesuai dengan Role untuk mengakses halaman masing-masing:
//$ curl -i -X POST -H "Content-Type: application/json" -d "{\"username\":\"user1\",\"password\":\"password1\"}" http://localhost:8000/signin

// Lebih baik melakukan test login dengan extensi chrome yaitu Talend API Tester: https://chrome.google.com/webstore/detail/talend-api-tester-free-ed/aejoelaoggembcahagimdiliamlcdmfm
// Untuk melakukan login dan mengakses halaman admin, student dan school, kita bisa menggunakan cookie yang sudah dibuat pada saat login

// Halaman admin bisa diakses oleh user dengan role admin
// Halaman student bisa diakses oleh user dengan role student
// Halaman school bisa diakses oleh user dengan role student dan admin
// Halaman public bisa diakses tanpa login
