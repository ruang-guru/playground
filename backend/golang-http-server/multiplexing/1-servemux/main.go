package main

import (
	"fmt"
	"log"
	"net/http"
)

// Pada implementasi handler sebelumnya, ketika kita hit ke endpoint manapun maka handler yang dieksekusi sama.
// Multiplexer berguna untuk melakukan route antara endpoint dengan handler
// Do golang, untuk mengimplementasikan multiplexer kita dapat menggunakan http.ServeMux

func main() {
	// Objek http.ServeMux yang bertindak sebagai multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home")
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	log.Println("Server run on port 8080")
	http.ListenAndServe("localhost:8080", mux)
}
