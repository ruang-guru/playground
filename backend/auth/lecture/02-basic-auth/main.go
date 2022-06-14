package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if user != "admin" && password != "123" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Write([]byte("Welcome Admin!"))
	})

	return mux
}

//Encode auth admin:123 disini -> https://www.base64encode.org/ atau menggunakan code sebelumnya yaitu 01-encode-base64

// Test menggunakan curl:
//$ curl -v -X GET http://localhost:8000/signin -H "Authorization: Basic YWRtaW46MTIz"
