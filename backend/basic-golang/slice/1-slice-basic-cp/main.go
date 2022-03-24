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
	var slice1 []string
	slice1 = append(slice1, "Yoga")
	slice1 = append(slice1, "Romadhon")
	fmt.Println(slice1)
	fmt.Print(len(slice1), cap(slice1))
=======
	//beginanswer
	var slice []string
	slice = append(slice, "Zein")
	slice = append(slice, "Fahrozi")
	fmt.Println(slice)
	fmt.Println(cap(slice), len(slice))
	//endanswer
>>>>>>> 2eae85288d82522736fd57b5cd8bd81e2676274f
}
