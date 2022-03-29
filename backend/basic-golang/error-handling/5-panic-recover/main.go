package main

import (
	"fmt"
	"time"
)

// Panic merupakan kondisi dimana program berhenti dan tidak dapat melanjutkan eksekusi program karena terdapat suatu error.
// Salah satu error yang dapat menyebabkan panic adalah pembagian dengan nol

func division(first, second int) int {
	return first / second
}

func divisionWithRecover(first, second int) int {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	return first / second
}

func main() {
	// Ketika kita tidak memberikan recover pada panic, maka program akan segera berhenti
	for {
		fmt.Println(division(1, 0))
		time.Sleep(time.Second * 1)
	}

	// Ketika kita memberikan recover pada panic, program akan terus berjalan (uncomment infinite loop dibawah dan comment infinite loop diatas untuk mencobanya!)
	// for {
	// 	fmt.Println(divisionWithRecover(1, 0))
	// 	time.Sleep(time.Second * 1)
	// }
}
