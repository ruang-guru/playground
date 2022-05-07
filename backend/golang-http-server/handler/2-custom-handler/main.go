package main

// Karena Handler merupakan sebuah interface dengan signature method ServeHTTP(ResponseWriter, *Request),
// maka kita dapat membuat custom handler.
// Untuk membuat custom handler, kita dapat membuat sebuah struct yang memiliki method ServeHTTP(ResponseWriter, *Request).

import (
	"fmt"
	"net/http"
	"time"
)

// Sebagai contoh, disini terdapat TimeHandler dengan atribut Timezone dan Format
type TimeHandler struct {
	Timezone string
	Format   string
}

// TimeHandler memiliki method ServeHTTP(ResponseWriter, *Request). Method ini mencetak hari, tanggal, bulan, dan tahun ke ResponseWriter.
func (h TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loc, err := time.LoadLocation(h.Timezone)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	t := time.Now().In(loc).Format(h.Format)
	w.Write([]byte(fmt.Sprintf("Today is %v in Timezone %s", t, h.Timezone)))
}

func main() {
	// objek dari TimeHandler dengan nilai Timezone Asia/Jakarta dan Format time.RFC3339
	handler := TimeHandler{
		Timezone: "Asia/Jakarta",
		Format:   time.RFC3339,
	}

	// Handler yang telah dibuat ditambahkan pada atribut Handler di Server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
