package main

import "fmt"

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
	var choice int
	const phi = 3.14
	var p, l, t, luas, s, a, r float64

	fmt.Println("Input:")
	fmt.Println("1. Rectange Area")
	fmt.Println("2. Rectangular Area")
	fmt.Println("3. Triangle Area")
	fmt.Println("4. Circle Area")
	fmt.Print("Enter Choice 1, 2, 3, or 4 : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Masukan Sisi : ")
		fmt.Scan(&s)
		luas = s * s
		fmt.Println("Luas Persegi adalah : ", luas)

	case 2:
		fmt.Print("Masukan Panjang : ")
		fmt.Scan(&p)

		fmt.Print("Masukan Lebar : ")
		fmt.Scan(&l)
		luas = p * l
		fmt.Println("Luas Persegi Panjang adalah : ", luas)

	case 3:
		fmt.Print("Masukkan panjang alas segitiga : ")
		fmt.Scan(&a)

		fmt.Print("Masukkan tinggi segitiga: ")
		fmt.Scan(&t)
		luas = 0.5 * a * t
		fmt.Println("Luas Persegi Panjang adalah : ", luas)

	case 4:
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&r)
		
		luas = phi * r * r
		fmt.Println("Luas lingkaran adalah : ", luas)

	default:
		fmt.Print("Invalid Choice")
	}

}
