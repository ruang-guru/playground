package main

import (
	"fmt"
)

// Check Point:
// Passing Grade
// - Input: Point
// - Output: Passing Grade Result

// Contoh:
// Input:
// - Masukkan nilai mahasiswa : 100 -> 0
// Output:
// - (=100) Lulus dengan nilai sempurna
// - (>70) Lulus
// - (=65) Hampir Lulus
// - (<70) Tidak lulus!. nilai anda: Point

func main() {

	var point int
	fmt.Printf("Masukkan nilai mahasiswa : ")
	fmt.Scan(&point)

	if point == 100 {
		fmt.Printf("Lulus dengan nilai sempurna")
	} else if point > 70 {
		fmt.Printf("Lulus")
	} else if point == 65 {
		fmt.Printf("Hampir Lulus")
	} else {
		fmt.Printf("Tidak lulus!. nilai anda: %d", point)
	}
}
