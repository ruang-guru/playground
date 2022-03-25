package main

import "fmt"

// Disini kalia coba melakukan if ketika kondisi nya sudah terpenuhi
// buatlah program yang akan menampilkan nilai dari setiap mahasiswa
// Jika nilai A = Cumlaude
// Jika nilai B = Lulus
// Jika nilai X = Tidak Lulus

// Kalian sudah tahukan cara mengakses value map pada Go meggunakan for-range kan?

func main() {
	name1 := "Zein Fahrozi"
	name2 := "Fabiansyah Raam"
	name3 := "Indra Kenz"

	mahasiswa := []map[string]string{
		{
			"name":  name1,
			"nilai": "A",
		},
		{
			"name":  name2,
			"nilai": "B",
		},
		{
			"name":  name3,
			"nilai": "X",
		},
	}

	// Output:
	/*
		Zein Fahrozi   Cumlaude
		Fabiansyah Raam   Lulus
		Indra Kenz   Tidak Lulus
	*/
	// TODO: answer here
	for _, v := range mahasiswa {
		if v["nilai"] == "A" {
			fmt.Println(v["name"], "cumlaude")

		} else if v["nilai"] == "B" {
			fmt.Println(v["name"], "Lulus")
		} else {
			fmt.Println(v["name"], "Tidak Lulus")
		}
		// fmt.Println(v["name"], v["nilai"])
	}
}
