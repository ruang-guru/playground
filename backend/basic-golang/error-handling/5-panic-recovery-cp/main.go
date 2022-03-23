package main

import "fmt"

func main() {
	// Dari contoh yang telah diberikan, coba kamu handle error panic code dibawah ini, sehingga code dapat terus berjalan sampai akhir.
	// Modifikasikan pada kode dari fungsi deferring()
	// Dan tampilkan error output yang ditangkap seperti ini:
	// 0. Books  is: The Eye of the World
	// 1. Books  is: The Great Hunt
	// 2. Books  is: The Dragon Reborn
	// Panic error terdeteksi: runtime error: index out of range [3] with length 3
	// Panic error terdeteksi: runtime error: index out of range [4] with length 3
	// Panic error terdeteksi: runtime error: index out of range [5] with length 3
	// Panic error terdeteksi: runtime error: index out of range [6] with length 3
	// Panic error terdeteksi: runtime error: index out of range [7] with length 3
	// Panic error terdeteksi: runtime error: index out of range [8] with length 3
	// Panic error terdeteksi: runtime error: index out of range [9] with length 3
	// We handled the panic

	for i := 0; i < 10; i++ {
		printBook(i)
	}
	fmt.Println("We handled the panic")
}

func printBook(i int) {
	// TODO: answer here

	books := []string{
		"The Eye of the World",
		"The Great Hunt",
		"The Dragon Reborn",
	}

	fmt.Printf("%v. Books  is: %v \n", i, books[i])
}

func deferring() {
	// TODO: answer here
}
