package main

import "fmt"

// Sekarang kita akan mencoba untuk membaca nilai yang kita telah inisialisasi.
// kita dapat membaca array dengan menggunakan index yang ada pada array tersebut.
func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Perlu diperhatikan perhitungan index array dimulai dari nol ya.
	fmt.Println(array[0]) // output: 1
	fmt.Println(array[9]) // output: 10
}
