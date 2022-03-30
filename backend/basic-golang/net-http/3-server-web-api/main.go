package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func viewTables(w http.ResponseWriter, r *http.Request) {
	// untuk menambahkan delimiter path
	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":  "Ruangguru Kampus Merdeka",
		"tables": data,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	http.HandleFunc("/view/tables", viewTables)

	fmt.Println("starting web server at http://localhost:8080/")
	// daftarkan server untuk listen di port 8080
	http.ListenAndServe(":8080", nil)
}

// how to test
//1. Akses browser ke URL: http://localhost:8080/view/tables
