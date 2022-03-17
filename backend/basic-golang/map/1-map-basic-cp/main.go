package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	var provinsi = map[string]string{
		"kalbar": "Pontianak",
		"kaltim": "Balikpapan",
	}

	for k, v := range provinsi {
		fmt.Printf(k, ":", v)
	}
}
