package main

import (
	"fmt"
	"math"
)

// Demo:
// Menghitung luas lingkaran
// - Input: jari-jari : nilai jari-jari lingkaran
// - Output: Jadi luasnya adalah : nilai luar lingkaran

// Contoh
// - Input: Masukkan sisi : 10
// - Output: Jadi luasnya adalah : 78.500000

func main() {
	var (
		r    float64
		area float64
	)
	const pi float64 = 3.14
	fmt.Printf("Masukkan jari-jari : ")
	fmt.Scan(&r)
	area = pi * (math.Pow(r, 2))
	fmt.Printf("Jadi luasnya adalah : %f", area)
}
