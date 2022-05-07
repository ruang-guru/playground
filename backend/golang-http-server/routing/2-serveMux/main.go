package main

import (
	"log"
	"net/http"
)

func main() {
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index"))
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", helloHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
