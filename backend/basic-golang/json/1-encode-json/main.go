package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Meja struct {
	Jenis  string `json:"jenis"`
	Warna  string
	Jumlah int `json:"jumlah"`
}

func main() {
	meja := Meja{
		Jenis:  "Meja Makan",
		Warna:  "Coklat",
		Jumlah: 20,
	}

	// kita bisa gunakan json.Marhal function untuk
	// encode variable meja ke JSON string
	mejaJSON, err := json.Marshal(meja)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}
	// mejaJSON adalah string JSON dalam format bytes
	// err adalah variable untuk menampung jika ada masalah
	// ketika proses encode

	//1. print JSON byte format tanpa casting ke string format
	fmt.Println("Print JSON tanpa casting ke string: ", mejaJSON)

	//2. Agar bisa terbaca kita perlu casting bytes ke bentuk string
	fmt.Println("\nPrint JSON dengan string casting: ", string(mejaJSON))

	//3. Perhatikan hasil print antara casting dan non casting
	//4. Perhatikan perbedaan antara field yang didefinisikan menggunakan json tag dan tidak
	//   pada field `Warna` tidak di definisikan json tag maka hasil encode key akan mengikuti nama field.
}
