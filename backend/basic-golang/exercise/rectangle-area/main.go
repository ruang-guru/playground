package main

import (
	"fmt"
)

// Demo:
// Menghitung luas persegi
// - Input: Masukkan sisi : nilai sisi
// - Output: luas persegi

// Contoh
// - Input: Masukkan sisi : 10
// - Output: Jadi luasnya adalah : 100.000000

func main() {
	var (
		size float32
		area float32
	)
	fmt.Printf("Masukkan sisi : ")
	fmt.Scan(&size)
	area = size * size
	fmt.Printf("Jadi luasnya adalah : %f", area)
}
