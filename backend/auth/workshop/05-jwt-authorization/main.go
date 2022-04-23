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
		// TODO: answer here
	})

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})

	return mux
}
