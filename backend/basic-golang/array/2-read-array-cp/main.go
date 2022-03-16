package main

import "fmt"

// Nah pada checkpoint kali ini kita akan membaca huruf pertama dan terakhir nama kalian
// Yuk dicoba, masih ingat kan dengan cara inisialisasi array pada contoh kode nomor 1 ?
// Outputkan hasilnya ya
func main() {
	// TODO: answer here
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Perlu diperhatikan perhitungan index array dimulai dari nol ya.
	fmt.Println(array[0]) // output: 1
	fmt.Println(array[9]) // output: 10
}
