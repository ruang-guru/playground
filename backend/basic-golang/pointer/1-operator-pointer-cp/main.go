package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.

// Print alamat memori dari masing-masing variabel dibawah ini
func main() {
	name := "Roger"
	age := 20
	isMarried := true

	//beginanswer
	fmt.Println(&name)
	fmt.Println(&age)
	fmt.Println(&isMarried)
	//endanswer
}
