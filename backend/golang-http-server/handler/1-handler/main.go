package main

// Untuk menambah kapabilitas dari server agar dapat menghandle request dari client,kita perlu membuat handler.
// Pada package net/http, terdapat interface Handler yang dapat digunakan untuk mengimplementasikan handler.
// Interface Handler memiliki satu method signature yaitu ServeHTTP(ResponseWriter, *Request).
// Salah satu implementasi dari interface Handler adalah HandlerFunc.

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Membuat handler dengan HandlerFunc
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}

	// Menambahkan handler pada objek server
	log.Println("Server run on port 8080")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
