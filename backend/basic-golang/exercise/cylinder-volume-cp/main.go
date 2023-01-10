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
	var (
		r   float32
		t   float32
		v   float32
		phi float32
	)
	fmt.Printf("Masukkan jari-jari: ")
	fmt.Scan(&r)
	fmt.Printf("Masukkan tinggi: ")
	fmt.Scan(&t)
	phi = 3.14
	// phi x r x r x t
	v = phi * r * r * t
	fmt.Println("Volume = ", v)
}
