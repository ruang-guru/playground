package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan empty interface.
// Cobalah untuk membuat sebuah data yang dapat merepresentasikan suatu menu pada restaurant.

func main() {
	var menu []map[string]interface{}

	// Buatlah beberapa menu yang memiliki atribut nama, jenis (nasi, cepat saji, minuman, dll), dan harga. Tambahkan setiap menu pada list `menu`
	//beginanswer
	ayamGoreng := make(map[string]interface{})
	ayamGoreng["Nama"] = "Ayam Goreng"
	ayamGoreng["Jenis"] = "Cepat saji"
	ayamGoreng["Harga"] = 20000

	menu = append(menu, ayamGoreng)

	cola := make(map[string]interface{})
	cola["Nama"] = "Cola"
	cola["Jenis"] = "Minuman"
	cola["Harga"] = 7000

	menu = append(menu, cola)
	//endanswer

	for _, m := range menu {
		for k, v := range m {
			fmt.Printf("%s: %v\n", k, v)
		}
	}
}
