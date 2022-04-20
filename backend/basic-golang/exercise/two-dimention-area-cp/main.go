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
	var sisi int
	var panjang int
	var lebar int
	var alas_segitiga int
	var tinggi_segitiga int
	var jari int
	var phi = 22 / 7

	var choice int = 0
	var result int = 0
	fmt.Println("1. Persegi")
	fmt.Println("2. Persegi Panjang")
	fmt.Println("3. Segitiga")
	fmt.Println("4. Lingkaran")
	fmt.Print("Hitung apa? ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Masukkan sisi: ")
		fmt.Scan(&sisi)
		result = sisi * sisi
		fmt.Print("Luas Persegi: ", result)

	case 2:
		fmt.Printf("Masukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Printf("Masukkan lebar: ")
		fmt.Scan(&lebar)
		result = panjang * lebar
		fmt.Print("Luas Persegi panjang: ", result)

	case 3:
		fmt.Printf("Masukkan alas: ")
		fmt.Scan(&alas_segitiga)
		fmt.Printf("Masukkan tinggi: ")
		fmt.Scan(&tinggi_segitiga)
		result = (alas_segitiga * tinggi_segitiga) / 2
		fmt.Print("Luas segitiga: ", result)

	case 4:
		fmt.Printf("Masukkan jari-jari: ")
		fmt.Scan(&jari)
		result = phi * jari * jari
		fmt.Print("Luas lingkaran: ", result)
	}

}
