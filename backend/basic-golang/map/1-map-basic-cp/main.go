package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	// TODO: answer here
	mapProvince := map[string]string{
		"Kalbar":  "Pontianak",
		"Kalteng": "Palangkaraya",
		"Kalsel":  "Banjarmasin",
		"Kaltim":  "Samarinda",
		"Kalut":   "Tanjung Selor",
	}

	for key, v := range mapProvince {
		fmt.Println("key:", key, "value: ", v)
	}
}
