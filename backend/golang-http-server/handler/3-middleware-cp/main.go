package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah sebuah middleware RequestMethodGetMiddleware yang berfungsi untuk mengecek request method.
// Jika request method tidak sama dengan GET, maka return HTTP status method not allowed

type TimeHandler struct {
	Timezone string
	Format   string
}

func (h TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loc, err := time.LoadLocation(h.Timezone)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	t := time.Now().In(loc).Format(h.Format)
	w.Write([]byte(fmt.Sprintf("Today is %v in Timezone %s", t, h.Timezone)))
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func main() {
	handler := TimeHandler{
		Timezone: "Asia/Jakarta",
		Format:   time.RFC3339,
	}

	log.Println("Server run on port 8080")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: RequestMethodGetMiddleware(handler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
