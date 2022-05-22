package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// encrypt token dari username menggunakan bcrypt kemudian kirim ke user kedalam cookie
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "token"
		var creds Credentials

		// Task:  Buat JSON body diconvert menjadi credential struct & return bad request ketika terjadi kesalahan decoding json

		// TODO: answer here
		checkErr := func(err error, httpError int) {
			if err != nil {
				w.WriteHeader(httpError)
				log.Fatal(err)
			}
		}
		bodyRead, err := ioutil.ReadAll(r.Body)
		checkErr(err, http.StatusBadRequest)

		err = json.Unmarshal(bodyRead, &creds)
		checkErr(err, http.StatusBadRequest)

		// Task: Ambil password dari username yang dipakai untuk login

		// TODO: answer here
		correctPassword := users[creds.Username]

		// Task: return unauthorized jika password salah

		// TODO: answer here
		if correctPassword != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Task: 1. Buat token string menggunakan bcrypt dari credential username
		// 		 2. return internal server error ketika terjadi kesalahan encrypting token

		// TODO: answer here
		token, err := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)
		checkErr(err, http.StatusInternalServerError)
		// Task: Set token string kedalam cookie response

		// TODO: answer here
		http.SetCookie(w, &http.Cookie{Name: cookieName, Value: string(token)})
	})

	return mux
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
