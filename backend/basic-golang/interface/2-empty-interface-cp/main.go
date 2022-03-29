package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan empty interface.
// Buatlah dua data makanan dan minuman yaitu ayam goreng dan cola yang memiliki atribut
// Nama, Jenis, dan Harga.
// Ayam Goreng, Cepat saji, 20000
// Cola, Minuman, 7000

func GetMenu() []map[string]interface{} {
	var menu []map[string]interface{}

	// TODO: answer here

	return menu
}

func main() {
	menu := GetMenu()

	for _, m := range menu {
		for k, v := range m {
			fmt.Printf("%s: %v\n", k, v)
		}
	}
}
