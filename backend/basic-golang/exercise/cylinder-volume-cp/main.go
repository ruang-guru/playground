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
	//beginanswer
	var (
		r      float32
		height float32
		volume float32
	)
	const pi = 3.14
	fmt.Printf("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)
	fmt.Printf("Masukkan tinggi tabung : ")
	fmt.Scan(&height)
	volume = pi * r * r * height
	fmt.Printf("Jadi volumenya adalah : %f\n", volume)
	//endanswer
}
