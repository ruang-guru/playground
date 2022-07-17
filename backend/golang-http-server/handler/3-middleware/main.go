package main

// Salah satu implementasi handler adalah middleware. Middleware dapat digunakan untuk mengecek
// suatu request dan melakukan action tertentu berdasarkan request yang diterima sebelum meneruskan ke handler selanjutnya.
// Middleware merupakan sebuah function yang memiliki parameter http.Handler dan memiliki nilai kembalian http.Handler.

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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

// Sebagai contoh, dibawah ini terdapat middleware getRequestMiddleware yang melakukan pengecekan request method dari suatu request.
// Jika request method tidak sama dengan GET, maka request tidak diteruskan ke handler selanjutnya.

func getRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Println("Method not allowed")
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
