package main

import "fmt"

// Nah pada checkpoint kali ini kita akan membaca huruf pertama dan terakhir nama kalian
// Yuk dicoba, masih ingat kan dengan cara inisialisasi array pada contoh kode nomor 1 ?
// Outputkan hasilnya ya
func main() {
	// TODO: answer here
	array := [4]string{"R", "A", "M", "A"}
	// Perlu diperhatikan perhitungan index array dimulai dari nol ya.
	fmt.Println(array[0]) // output: 1
	fmt.Println(array[3])
}
