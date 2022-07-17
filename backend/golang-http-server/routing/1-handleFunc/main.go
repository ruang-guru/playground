package main

import (
	"log"
	"net/http"
)

// Salah satu implementasi sederhana dari routing di golang adalah dengan menggunakan http.HandleFunc.
// Di http.HandleFunc terdapat dua parameter yaitu route dan handler.

func main() {
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index"))
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	// Route "/" dengan indexHandler
	http.HandleFunc("/", indexHandler)

	// Route "/hello" dengan helloHandler
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
