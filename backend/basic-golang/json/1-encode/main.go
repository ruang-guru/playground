package main

import (
	"encoding/json"
	"fmt"
)

// Go memiliki pendekatan yang unique untuk mengimplementasi JSON data.
// Salah satu cara yang biasa digunakan adalah dengan encoding JSON kedalam `struct`.
// Ketika kita encoding / decoding `struct` ke JSON, defaultnya akan menggunakan nama field
// didalam `struct` kecuali kita definisikan JSON tag di setiap field.

type User struct {
	FirstName string `json:"firstName"` // key akan menjadi "firstName"
	BirthYear int    `json:"birtYear"`  // key akan menjadi "birthYear"
	Email     string // key akan menjadi "Email"
	Address   string `json:"address,omitempty"`
	Hobby     string `json:"-"`
}

// Untuk encode json kita bisa menggunakan function `Marshal`.

func main() {
	jsonData, _ := json.Marshal(User{
		FirstName: "Rogu",
		BirthYear: 2000,
		Email:     "example@ruangguru.com",
		Address:   "",
		Hobby:     "belajar",
	})
	jsonString := string(jsonData)
	fmt.Println(jsonString)
}
