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
	const PI = 3.14
	var (
		choice int
		area   float32
		shape  string
	)
	fmt.Println("Input :")
	fmt.Println("1: Rectangle Area")
	fmt.Println("2: Rectangular Area")
	fmt.Println("3: Triangle Area")
	fmt.Println("4: Circle Area")
	fmt.Print("Enter choice 1, 2, 3, or 4: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var sisi float32
		shape = "Persegi"
		fmt.Println("(1) Rectange Area")
		fmt.Print("Masukkan sisi: ")
		fmt.Scan(&sisi)
		area = float32(math.Pow(float64(sisi), 2))

	case 2:
		var panjang, lebar float32
		shape = "Persegi Panjang"
		fmt.Println("(1) Rectangular Area")
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&lebar)
		area = panjang * lebar

	case 3:
		var alas, tinggi float32
		shape = "Segitiga"
		fmt.Println("(3) Triangle Area")
		fmt.Print("Masukkan panjang alas segitiga: ")
		fmt.Scan(&alas)
		fmt.Print("Masukkan tinggi: ")
		fmt.Scan(&tinggi)
		area = alas * tinggi / 2

	case 4:
		var r float32
		shape = "Lingkaran"
		fmt.Println("(4) Circle Area")
		fmt.Print("Masukkan jari-jari: ")
		fmt.Scan(&r)
		area = PI * r * r

	default:
		fmt.Print("Invalid Input")
		return
	}

	fmt.Printf("(%d) Luas %s adalah %.6f", choice, shape, area)
}
