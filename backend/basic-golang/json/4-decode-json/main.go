package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// encode JSON string nested object berikut
/*
{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
		}
}
*/

type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Meja struct {
	Jenis  string `json:"jenis"`
	Warna  string `json:"warna"`
	Jumlah int    `json:"jumlah"`
	Ukuran Ukuran `json:"ukuran"`
}

func main() {
	mejaJson := `{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
		}
	}`

	var meja Meja
	if err := json.Unmarshal([]byte(mejaJson), &meja); err != nil {
		log.Fatalf("JSON Unmarshal error: %v", err)
	}

	// Print semua value dari object meja
	fmt.Println("print semua: ", meja)

	// Print specifict value dari meja
	fmt.Println("jenis meja: ", meja.Jenis)
	fmt.Println("warna meja: ", meja.Warna)
}
