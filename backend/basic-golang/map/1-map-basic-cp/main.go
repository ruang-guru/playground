package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	// TODO: answer here
	var provinsi = map[string]string{
		"Kalbar": "Pontianak",
		"Kaltim": "Balikpapan",
		"Jabar":  "Bandung",
		"Bali":   "Denpasar",
	}

	for k, v := range provinsi {
		fmt.Println(k, ":", v)
	}
}
