package main

import "fmt"

// Disini teman teman akan mencoba untuk
// melakukan penambahan data pada slice.
// Buatlah variable slice dengan tipe data string.
// Lalu masukan nama kalian ke dalam slice.
// Expected outout: ["NamaPanggilan", "Nama Akhir"]
// Contoh [Zein Fahrozi]
// Outputkan jawabannya ya pastikan cap dan len nya adalah 2
func main() {
	//beginanswer
	var slice []string
	slice = append(slice, "Zein")
	slice = append(slice, "Fahrozi")
	fmt.Println(slice)
	fmt.Println(cap(slice), len(slice))
	//endanswer
}
