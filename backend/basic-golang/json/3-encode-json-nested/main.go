package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// buat struct untuk nanti dijadikan type data/nested di struct meja
type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Meja struct {
	Jenis  string `json:"jenis"`
	Warna  string `json:"warna"`
	Jumlah int    `json:"jumlah"`
	// assign field name `Ukuran` dengan
	// struct `Ukuran` yang sudah dibuat sebelumnya
	Ukuran Ukuran `json:"ukuran"`
}

func main() {
	// set value kedalam struct yang sudah dibuat
	meja := Meja{
		Jenis:  "Meja Makan",
		Warna:  "Coklat",
		Jumlah: 20,
		Ukuran: Ukuran{
			Panjang: "50 cm",
			Tinggi:  "25 cm",
		},
	}

	mejaJSON, err := json.Marshal(meja)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}

	fmt.Println(string(mejaJSON))

}
