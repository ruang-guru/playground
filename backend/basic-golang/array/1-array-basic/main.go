package main

import "fmt"

// Kalian bisa liat lebih detail disini https://go101.org/article/container.html
func main() {
	/*
		ada beberapa cara melakukan inisiallisasi array dalam Go
		var namaArray [size]type

		- var array1 [10]int	<- Data kosong
		- array1 := [10]int{1,2,3,4,5,6,7,8,9,10} <- Data ada
	*/

	var array1 [10]int
	fmt.Println(array1) // output: [0 0 0 0 0 0 0 0 0 0]

	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array2)
}
