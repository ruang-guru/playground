package main

import "fmt"

// Ada tiga jenis logical operator dalam Go

/*
	&& (and)
	|| (or)
	! (not)
*/

func main() {
	angka1 := 1
	angka2 := 2
	angka2lagi := 2
	angka3 := 3

	// angka1 == angka2 := false . not false = true
	if !(angka1 == angka2) {
		fmt.Println("angka 1 tidak sama dengan angka 2")
	}

	// angka2 == angka2lagi = true, angka1 == angka 3 = false
	// true or false = true
	if angka2 == angka2lagi || angka1 == angka3 {
		fmt.Println("angka 2 sama dengan angka 2 lagi")
	}
}
