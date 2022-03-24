package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan empty interface.
// Cobalah untuk membuat sebuah data yang dapat merepresentasikan suatu menu pada restaurant.

func main() {
	var menu []map[string]interface{}

	// Buatlah beberapa menu yang memiliki atribut nama, jenis (nasi, cepat saji, minuman, dll), dan harga. Tambahkan setiap menu pada list `menu`
	// TODO: answer here

	for _, m := range menu {
		for k, v := range m {
			fmt.Printf("%s: %v\n", k, v)
		}
	}
}
