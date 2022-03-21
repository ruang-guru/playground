package main

import (
	"fmt"
	"math"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here
	var choice int = 0
	var result float32

	fmt.Println("1: Rectange Area")
	fmt.Println("2: Rectangular Area")
	fmt.Println("3: Triangle Area")
	fmt.Println("4: Circle Area")

	fmt.Println("Enter choice 1, 2, 3, or 4:")
	fmt.Scan(&choice)

	switch choice {
		case 1:
			var side float32
			fmt.Println("Masukan sisi:")
			fmt.Scan(&side)
			
			result = side * side
			fmt.Println("Luas Persegi adalah:", result)
		case 2:
			var length, wide float32
			fmt.Println("Masukkan panjang:")
			fmt.Scan(&length)
			fmt.Println("Masukkan lebar:")
			fmt.Scan(&wide)

			result = length * wide
			fmt.Println("Luas Persegi Panjang adalah:", result)
		case 3:
			var a, t float32
			fmt.Println("Masukkan panjang alas segitiga:")
			fmt.Scan(&a)
			fmt.Println("Masukkan tinggi segitiga:")
			fmt.Scan(&t)

			result = a * t * .5
			fmt.Println("Luas Segitiga adalah:", result)
		case 4:
			var r, res float64
			const pi float64 = 3.14
			fmt.Println("Masukkan jari-jari:")
			fmt.Scan(&r)

			res = pi * (math.Pow(r, 2))
			fmt.Println("Luas Lingkaran adalah:", res)
		default:
			fmt.Println("Invalid value")
	}
}
