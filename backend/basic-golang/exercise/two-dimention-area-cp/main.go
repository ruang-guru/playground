package main

import (
	"fmt"
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
	var (
		s    float32
		p    float32
		l    float32
		t    float32
		r    float32
		a    float32
		nama string
	)

	var choice int = 0
	var result float32
	const pi = 3.14

	fmt.Println("1. Rectange Area\n2.Rectangular Area\n3. Triangle Area\n4. Circle Area")
	fmt.Printf("Enter Choice : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Input sisi : ")
		fmt.Scan(&s)
		result = s * s
		nama = "Persegi"
	case 2:
		fmt.Printf("Input panjang : ")
		fmt.Scan(&p)
		fmt.Println("Input lebar : ")
		fmt.Scan(&l)
		result = p * l
		nama = "Persegi Panjang"
	case 3:
		fmt.Println("Input alas segitiga : ")
		fmt.Scan(&a)
		fmt.Println("Input tinggi: ")
		fmt.Scan(&t)
		result = 1 / 2 * a * t
		nama = "segitiga"
	case 4:
		fmt.Println("Input jari-jari : ")
		fmt.Scan(&r)
		result = pi * r * r
		fmt.Println("Luas")
		nama = "lingkaran"
	default:
		fmt.Println("Invalid Choice")
	}

	fmt.Printf("Luas %s adalah %0.2f\n", nama, result)

}
