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
		// JSON body diconvert menjadi creditial struct & return bad request ketika terjadi kesalahan decoding json

		// TODO: answer here

		// Task: Ambil password dari username yang dipakai untuk login & return unauthorized jika password salah

		// TODO: answer here

		// Task: 1. Deklarasi expiry time untuk token jwt
		// 		 2. Buat claim menggunakan variable yang sudah didefinisikan diatas
		//       3. Expiry time menggunakan time millisecond

		// TODO: answer here

		// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai

		// TODO: answer here

		// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		//       2. Buat return internal error ketika ada kesalahan ketika pembuatan JWT string

		// TODO: answer here

		// Task: Set token string kedalam cookie response

		// TODO: answer here
	})

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		// TODO: answer here

		// Task: Ambil value dari cookie token

		// TODO: answer here

		// Task: Deklarasi variable claim

		// TODO: answer here

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada

		// TODO: answer here

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here

		// Task: return unauthorized ketika role user tidak sesuai dengan role admin

		// TODO: answer here

		// Task: return data dalam claim, seperti username yang telah didefinisikan

		// TODO: answer here
	})

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		// TODO: answer here

		// Task: Ambil value dari cookie token

		// TODO: answer here

		// Task: Deklarasi variable claim

		// TODO: answer here

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada

		// TODO: answer here

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here

		// Task: return data dalam claim, seperti username yang telah didefinisikan

		// TODO: answer here
	})

	return mux
}
