package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Dari contoh sebelumnya tambahkan logic
// untuk handle POST request add bulk table
// endpoint URL: http://localhost:8080/tables
// request body:
/*
[
    {
        "id": "T012",
        "type": "Meja Rias",
        "color": "Coklat",
        "total": 3
    },
    {
        "id": "T011",
        "type": "Meja Kotak",
        "color": "Hitam",
        "total": 4
    }
]
*/

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TablesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// validate methode request
	// logic handler untuk GET request
	if r.Method == "GET" {
		// encode data ke dalam format string JSON
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// untuk mendaftarkan result sebagai response
		w.Write(result)
		return
	}

	// logic handle POST request
	if r.Method == "POST" {
		// TODO: answer here

		// set header response code with status created/201
		w.WriteHeader(http.StatusCreated)
		// write json reponse body
		w.Write([]byte(`{"status":"add tables succeed"}`))
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
	http.HandleFunc("/tables", TablesHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	// daftarkan server untuk listen di port 8080
	http.ListenAndServe(":8080", nil)
}
