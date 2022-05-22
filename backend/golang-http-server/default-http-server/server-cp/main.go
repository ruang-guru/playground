package main

import (
	"log"
	"net/http"
)

// Dari contoh yang diberikan, buatlah implementasi server dengan menggunakan struct Server dari package http/net/
// Buatlah server dengan address localhost dan di port 3000

func main() {
	// TODO: answer here
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
