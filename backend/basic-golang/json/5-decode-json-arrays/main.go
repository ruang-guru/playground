package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// encode Array JSON string nested object berikut
/*
[
	{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
			"panjang": "70 cm",
			"tinggi": "30 cm"
		}
	},
	{
		"jenis": "Meja Makan",
		"warna": "Coklat",
		"jumlah": 1,
		"ukuran": {
			"panjang": "30 cm",
			"tinggi": "20 cm"
		}
	}
]
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
	mejaJson := `[
		{
			"jenis": "Meja Lipat",
			"warna": "Hitam",
			"jumlah": 1,
			"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
			}
		},
		{
			"jenis": "Meja Makan",
			"warna": "Coklat",
			"jumlah": 1,
			"ukuran": {
				"panjang": "30 cm",
				"tinggi": "20 cm"
			}
		}
	]`

	var mejaMeja []Meja
	if err := json.Unmarshal([]byte(mejaJson), &mejaMeja); err != nil {
		log.Fatalf("JSON Unmarshal error: %v", err)
	}

	// Print semua value dari object meja
	fmt.Println("print semua: ", mejaMeja)

	// Print specifict value dari meja
	// untuk array bisa ada dua cara untuk
	// ambil specifict value
	// 1. dengan looping array object
	// 2. print dengan specifict index

	// cara pertama get value dengan loop
	fmt.Println("\n\nget value dengan loop:")
	for _, meja := range mejaMeja {
		fmt.Println("jenis meja: ", meja.Jenis)
		fmt.Println("warna meja: ", meja.Warna)
		fmt.Println("tinggi meja: ", meja.Ukuran.Tinggi)
	}

	// cara kedua dengan specifict index
	fmt.Println("\n\nget value dengan index:")
	fmt.Println("jenis meja: ", mejaMeja[0].Jenis)
	fmt.Println("warna meja: ", mejaMeja[0].Warna)
	fmt.Println("jenis meja: ", mejaMeja[1].Jenis)
	fmt.Println("warna meja: ", mejaMeja[1].Warna)

	// bandingkan kedua cara tersebut mana yang lebih efisien
}
