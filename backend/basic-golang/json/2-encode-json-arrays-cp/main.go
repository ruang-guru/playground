package main

import (
	"encoding/json"
	"log"
)

// Buat string JSON dengan hasil seperti berikut
// [{"jenis":"Meja Lipat","warna":"Coklat","jumlah":40,"deskripsi":"meja untuk belajar"},{"jenis":"Meja Hijau","warna":"Hijau","jumlah":10,"deskripsi":"meja untuk pengadilan"}]

type Meja struct {
	// TODO: answer here
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
}

func NewMeja(m Items) Items {
	return m
}
