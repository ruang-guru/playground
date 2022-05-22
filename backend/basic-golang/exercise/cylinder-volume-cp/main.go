package main

import "fmt"

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
	const PI = 3.14
	var (
		r      float32
		t      float32
		volume float32
	)
	fmt.Printf("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Printf("Masukkan tinggi tabung : ")
	fmt.Scan(&t)

	volume = PI * r * r * t
	fmt.Printf("Jadi volumenya adalah : %.6f", volume)
}
