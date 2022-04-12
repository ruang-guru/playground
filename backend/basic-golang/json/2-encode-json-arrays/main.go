package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Meja struct {
	Jenis  string `json:"jenis"`
	Warna  string `json:"warna"`
	Jumlah int    `json:"jumlah"`
}

func main() {
	// Buat struct array dengan menambahka bracket `[]`
	// di depan nama struct
	mejaMeja := []Meja{
		{
			Jenis:  "Meja Makan",
			Warna:  "Coklat",
			Jumlah: 20,
		},
		{
			Jenis:  "Meja Belajar",
			Warna:  "Hijau",
			Jumlah: 2,
		},
		{
			Jenis:  "Meja Kerja",
			Warna:  "Hitam",
			Jumlah: 1,
		},
	}

	// kita bisa gunakan json.Marhal function untuk
	// encode variable mejaMeja ke JSON string
	mejaMejaJSON, err := json.Marshal(mejaMeja)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}
	// mejaMejaJSON adalah string JSON dalam format bytes
	// err adalah variable untuk menampung jika ada masalah
	// ketika proses encode

	fmt.Println(string(mejaMejaJSON))
}

//notes:
//perhatikan bracket untuk menuliskan struct dalam format array
//[] untuk mendefinisikan type data array
//ketika inisiasi value kedalam array struct
//perlu menambahkan {} sebelum fieldname
