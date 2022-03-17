package main

import "fmt"

// Didalam array kita bisa menggunakan fungsi `len` dan `cap`
// untuk mengetahui panjang dan kapasitas array tersebut.
// perbedaan keduanya adalah len adalah panjang array dan cap adalah kapasitas array.
func main() {
	var array1 [10]int
	fmt.Println(len(array1), cap(array1)) // Output: 10 10

	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(len(array2), cap(array2)) // Output: 10 10

	// Nah kalau diperhatikan kedua nya sama ya.
	// Nah bagaimana klo seperti ini

	var array3 []int
	fmt.Println(len(array3), cap(array3)) // Output: 0 0
	// Hmm, 0, 0 ya. karena kita tidak memiliki array. Tetapi sebuah slice.
	// Tenang kita akan belajar slice selanjutnya.

	// Coba kita isi slice tersebut
	array3 = append(array3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(array3), cap(array3)) // Output: 7, 8
	// Kok kita gk bisa masukin data dengan cara  array3[0] = 1 ya ? Nah akan dijelaskan di materi slice ya.
}
