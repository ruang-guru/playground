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
<<<<<<< HEAD
	// TODO: answer here
	// nama := []string{"reza", "wijaya"}
	// fmt.Println(nama)
	nama := []string{}
	nama = append(nama, "reza")
	nama = append(nama, "wijaya")
	fmt.Println(nama)
	fmt.Println(cap(nama), len(nama))

=======
	//beginanswer
	var slice []string
	slice = append(slice, "Zein")
	slice = append(slice, "Fahrozi")
	fmt.Println(slice)
	fmt.Println(cap(slice), len(slice))
	//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
}
