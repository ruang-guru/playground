package main

import "fmt"

// Empty interface dapat menyimpan berbagai tipe data.
// Empty interface biasanya digunakan ketika terdapat kemungkinan lebih dari satu tipe data yang disimpan pada variabel.

func main() {
	var data interface{}
	data = 100
	fmt.Println(data)

	data = "Hello world"
	fmt.Println(data)

	data = true
	fmt.Println(data)

	// Empty interface dapat digunakan untuk membuat map dengan tipe data pada value lebih dari satu jenis tipe data
	person := make(map[string]interface{})
	person["Name"] = "Roger"
	person["Age"] = 50
	person["Nationality"] = "USA"
	person["IsMarried"] = true

	fmt.Println(person["Name"])
	fmt.Println(person["Age"])
	fmt.Println(person["Nationality"])
	fmt.Println(person["IsMarried"])
}
