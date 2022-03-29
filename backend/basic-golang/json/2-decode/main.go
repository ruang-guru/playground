package main

import (
	"encoding/json"
	"fmt"
)

// Decode dapat dilakukan untuk mengubah json menjadi objek dari sebuah struct.
// Untuk decode kita bisa gunakan function `Unmarshal`

type User struct {
	FirstName string `json:"firstName"` // key akan menjadi "firstName"
	BirthYear int    `json:"birtYear"`  // key akan menjadi "birthYear"
	Email     string // key akan menjadi "Email"
	Address   string `json:"address,omitempty"`
	Hobby     string `json:"-"`
}

func main() {
	byteJSONData := []byte(`{
		"firstName":"Rogu",
		"birthYear":2000,
		"Email":"example@ruangguru.com",
		"address": "jl. ruangguru",
		"Hobby": "belajar"
	}`)

	u := User{}
	err := json.Unmarshal(byteJSONData, &u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u)
}
