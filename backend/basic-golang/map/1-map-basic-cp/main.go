package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	// TODO: answer here
	provinsi := map[string]string{"kaltim": "samarinda",
		"kalteng": "palangkaraya",
		"kalsel":  "banjarmasin"}

	// untuk mengakses definisi (value) kita harus mengetahui key nya

	fmt.Println(provinsi["kaltim"])
	fmt.Println(provinsi["kalteng"])
	fmt.Println(provinsi["kalsel"])
}
