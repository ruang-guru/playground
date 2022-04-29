package main

// Dalam golang, untuk membuat HTTP server kita dapat menggunakan package net/http.
// Pada package net/http terdapat struct Server yang dapat digunakan untuk mengimplementasikan server HTTP.

import (
	"log"
	"net/http"
)

func main() {
	// server http dengan address localhost:8080
	server := http.Server{
		Addr: "localhost:8080",
	}

	log.Println("Server running on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// cara lain yang dapat digunakan adalah dengan menggunakan http.ListenAndServe
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
