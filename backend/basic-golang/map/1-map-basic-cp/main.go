package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
<<<<<<< HEAD
	// TODO: answer here
	provinsi := map[string]string{
		"Kalimantan Barat":   "Pontianak",
		"Kalimantan Tengah":  "Palangka Raya",
		"Kalimantan Selatan": "Banjarmasin",
		"Kalimantan Timur":   "Samarinda",
	}

	for key, item := range provinsi {
		fmt.Println(key, ":", item)
	}
=======
	//beginanswer
	var provinsi = map[string]string{
		"Kalimantan Barat":   "Pontianak",
		"Kalimantan Tengah":  "Palangkaraya",
		"Kalimantan Selatan": "Samarinda",
		"Kalimantan Timur":   "Balikpapan",
		"Kalimantan Utara":   "Tanjung Selor",
	}

	for k, v := range provinsi {
		fmt.Println(k, ":", v)
	}
	//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
}
