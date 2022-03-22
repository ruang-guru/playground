package main

import "fmt"

// pada cekpoin ini kalian akan mencoba untuk menghapus data dengan tipe []map[string]string
// gabungan slice dan map.

func main() {
	var namaUmur = []map[string]string{
		{"name": "Socrates", "gender": "male"},
		{"name": "Plato", "gender": "male"},
		{"name": "Ada", "gender": "female"},
		{"name": "Leonhard Euler", "gender": "female"},
		{"name": "Blaise Pascal", "gender": "male"},
	}

	// terdapat kesalahan pada data gender tersebut dapatkan kalian memperbaiki nya ?
<<<<<<< HEAD
	// TODO: answer here

	fmt.Println("=========== CHANGE GENDER ===========")
	namaUmur[3]["gender"] = "male"
=======
	//beginanswer
	namaUmur[3]["gender"] = "male"
	//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	for _, val := range namaUmur {
		fmt.Println(val["name"], " ", val["gender"])
	}

	// Nah coba saatnya kalian menghapuskan key "gender" pada setiap data
	// delete data if key is equal "gender"
	fmt.Println("=========== DELETE GENDER ===========")
	for _, val := range namaUmur {
		if val["gender"] == "male" {
			delete(val, "gender")
		}
		fmt.Println(val)
	}
	// Output sebelum dihapus
	/*
		map[gender:male name:Socrates]
		map[gender:male name:Plato]
		map[gender:female name:Ada]
		map[gender:male name:Leonhard Euler]
		map[gender:male name:Blaise Pascal]
	*/

	//beginanswer
	for _, val := range namaUmur {
		delete(val, "gender")
	}
	//endanswer

	// Output setelah dihapus
	/*
		map[name:Socrates]
		map[name:Plato]
		map[name:Ada]
		map[name:Leonhard Euler]
		map[name:Blaise Pascal]
	*/
}
