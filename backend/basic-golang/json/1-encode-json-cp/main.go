package main

import (
	"encoding/json"
	"log"
)

// Dari struct dan nama field yang sama dari contoh
// Buat string JSON dengan hasil seperti berikut
// {"jenis":"Meja Belajar","color":"green","jumlah":2}

type Meja struct {
	// TODO: answer here
}

func (m Meja) EncodeJSON() string {
	// TODO: answer here
}

func NewMeja(m Meja) Meja {
	return m
}
