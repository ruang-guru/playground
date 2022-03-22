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
	fmt.Println("===========Two Dimention Area=========")
	fmt.Println("1: Rectange Area")
	fmt.Println("2: Rectangular Area")
	fmt.Println("3: Triangle Area")
	fmt.Println("4: Circle Area")
	fmt.Print("Enter choice 1, 2, 3, or 4:")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var side float32
		fmt.Print("Masukkan sisi: ")
		fmt.Scan(&side)
		result = side * side
		fmt.Printf("Luas persegi adalah: %f\n", result)
	case 2:
		var length, width float32
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&length)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&width)
		result = length * width
		fmt.Printf("Luas persegi panjang adalah: %f\n", result)
	case 3:
		var a, t float32
		fmt.Print("Masukkan alas segitiga: ")
		fmt.Scan(&a)
		fmt.Print("Masukkan tinggi segitiga: ")
		fmt.Scan(&t)
		result = 0.5 * a * t
		fmt.Printf("Luas segitiga adalah: %f\n", result)
	case 4:
		var r, res float64
		const pi float64 = 3.14
		fmt.Print("Masukkan jari jari: ")
		fmt.Scan(&r)
		res = pi * (math.Pow(r, 2))
		fmt.Printf("Luas segitiga adalah: %f\n", res)
	}
}
