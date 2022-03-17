package main

import "fmt"

// Nah pada checkpoint kali ini kita akan membaca huruf pertama dan terakhir nama kalian
// Yuk dicoba, masih ingat kan dengan cara inisialisasi array pada contoh kode nomor 1 ?
// Outputkan hasilnya ya
func main() {
	// TODO: answer here
	array1 := [4]string{"Y", "o", "g", "a"}
	fmt.Println(array1[0])
	fmt.Println(array1[len(array1)-1])
}
