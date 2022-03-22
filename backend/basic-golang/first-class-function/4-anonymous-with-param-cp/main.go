package main

import "fmt"

func main() {
	//fungsi yang mengembalikan nilai pangkat dua dari parameter yang diterima
	//contoh input parameter yang diterima (3)
	//maka fungsi akan mengembalikan 9

	//beginanswer
	square := func(angka int) int {
		return angka * angka
	}(5)
	//endanswer
	fmt.Println(square)
}
