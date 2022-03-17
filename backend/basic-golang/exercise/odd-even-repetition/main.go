package main

import (
	"fmt"
)

// Demo:
// Odd Even Repetition
// - Input: Masukkan jumlah angka : nilai
// - Output: Deretan nilai dengan tanda ganjil atau genap yang sesuai

// Contoh
// - Input: Masukkan jumlah angka : 5
// - Output:
// 1 ganjil
// 2 genap
// 3 ganjil
// 4 genap
// 5 ganjil

func main() {
	var size int
	fmt.Printf("Masukkan jumlah angka : ")
	fmt.Scan(&size)

	for i := 1; i <= size; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ganjil\n", i)
		} else {
			fmt.Printf("%d genap\n", i)
		}
	}
}
