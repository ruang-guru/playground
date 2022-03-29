package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func tables(w http.ResponseWriter, r *http.Request) {
	// set header content-type dengan json untuk menentukan response type
	// dalam contoh ini akan response dengan format json
	w.Header().Set("Content-Type", "application/json")

	// validate methode request
	if r.Method == "GET" {
		// encode data ke dalam format string JSON
		result, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// untuk mendaftarkan result sebagai response
		w.Write(result)
		return
	}

	// mengembalikan response error bad request
	// jika request method tidak sesuai
	http.Error(w, "", http.StatusBadRequest)
}

func table(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// untuk mengambil value yang dikirim oleh client dengan key `id`
		id := r.FormValue("id")

		// validate table id yang sesuai dengan id yang dikirim client
		for _, table := range data {
			if table.ID == id {
				result, err := json.Marshal(table)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// daftarkan response `result` jika ditemukan id yang sesuai dengan input client
				w.Write(result)
				return
			}
		}

		// kembalikan response not found jika data yang diinput client tidak ditemukan
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	// daftarkan routing endpoint yang dapat dilayani oleh server
	// parameter 1 adalah endpoint route,
	// parameter 2 adalah function dimana kita melakukan prosesing data untuk response
	http.HandleFunc("/table", table)
	http.HandleFunc("/tables", tables)

	fmt.Println("starting web server at http://localhost:8080/")
	// daftarkan server untuk listen di port 8080
	http.ListenAndServe(":8080", nil)
}

//How to test:
//1. Akses melalui browser ke endpoint yang tersedia
//	- untuk menampilkan semua data table akses ke http://localhost:8080/tables
//  - untuk menampilkan table id tertentu akses ke http://localhost:8080/table?id=T002
//2. menggunakan curl melalui commandline
//	- curl http://localhost:8080/tables
//3. menggunakan postman
