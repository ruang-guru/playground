package main

import "fmt"

// Pointer di golang memiliki dua operator yaitu address operator dan indirect operator.
// Address operator (&) digunakan untuk mendapatkan alamat memori dari suatu variabel.
// Indirect operator (*) digunakan untuk mendapatkan nilai dari alamat memori yang ditunjuk.

func main() {
	var score int = 100

	// Untuk mendapatkan alamat memori kita dapat menggunakan address operator
	scoreMemoryAddress := &score
	fmt.Printf("Alamat memory dari variabel score adalah %v\n", scoreMemoryAddress)

	// Untuk mendapatkan nilai dari alamat memori tertentu, kita dapat menggunakan indirect operator
	fmt.Printf("Nilai dari alamat memory %v adalah %d\n", scoreMemoryAddress, *scoreMemoryAddress)
}
