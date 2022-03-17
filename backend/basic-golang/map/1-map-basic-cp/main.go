package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	m0 := map[string]string{
		"Kalimantan Barat":   "Pontianak",
		"Kalimantan Timur":   "Samarinda",
		"Kalimantan Selatan": "Banjarmasin",
		"Kalimantan Tengah":  "Palangkaraya",
		"Kalimantan Utara":   "Tanjung Selor",
	}

	for key, val := range m0 {
		fmt.Println(key, val)
	}
}
