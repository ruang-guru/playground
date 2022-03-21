package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
	var r, t, v float64
	const phi = 3.14
	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)
	fmt.Print("Masukkan tinggi tabung : ")
	fmt.Scan(&t)
	v = phi * r * r * t
	fmt.Printf("Jadi volumenya adalah : %f\n	", v)

}
