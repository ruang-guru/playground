package main

import (
	"fmt"
)

// Check Point:
// Passing Grade
// - Input: Point
// - Output: Status Error or Passing Grade Result

// Contoh:
// Input:
// - Masukkan nilai mahasiswa : 100 -> 0
// Output:
// - (>100) Nilai tidak boleh lebih dari 100
// - (<0) Nilai tidak boleh kurang dari 0
// - (=100) Lulus dengan nilai sempurna
// - (>70) Lulus
// - (=65) Hampir Lulus
// - (<70) Tidak lulus!. nilai anda: Point

func main() {
	var point int
	fmt.Printf("Masukkan nilai mahasiswa : ")
	fmt.Scan(&point)

	// TODO: answer here

}
