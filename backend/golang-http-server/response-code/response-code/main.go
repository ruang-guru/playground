package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Pada package net/http golang terdapat const yang berisi HTTP response code.
// Untuk const HTTP response code di golang, bisa kamu akses link berikut https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/net/http/status.go;l=9

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

// Sebagai contoh, dibawah ini pada middleware getRequestMiddleware, jika request method tidak sama dengan GET
// maka akan return response code 405 dimana response code tersebut adalah Method Not Allowed

func getRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method is not allowed.", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := TimeHandler{
		Timezone: "Asia/Jakarta",
		Format:   time.RFC3339,
	}

	log.Println("Server run on port 8080")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: getRequestMiddleware(handler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
