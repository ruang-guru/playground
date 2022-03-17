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
	var msg string
	fmt.Printf("Masukkan nilai mahasiswa : ")
	fmt.Scan(&point)

	// TODO: answer here
	if point > 100 {
		msg = "Nilai tidak boleh lebih dari 100"
	} else if point < 0 {
		msg = "Nilai tidak bolehh kurang dari 0"
	} else {

		if point == 100 {
			msg = "Lulus dengan nilai sempurna"
		} else if point > 70 {
			msg = "Lulus"
		} else if point == 65 {
			msg = "Hampir lulus"
		} else {
			msg = ("Tidak lulus! nilai anda : %d")
		}
	}
	fmt.Printf(msg, point)

}
