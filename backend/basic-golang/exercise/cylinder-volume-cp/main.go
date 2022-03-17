package main

import (
	"fmt"
	"math"
	
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
	var in float64
	var inH float64
	fmt.Scan(&in)
	fmt.Scan(&inH)
	var result = math.Pi*in*in*inH
	fmt.Println(result)


}
