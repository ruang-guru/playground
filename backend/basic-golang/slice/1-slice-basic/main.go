package main

import "fmt"

/*
slice types: []T
Nah kalian tau kalau slice itu adalah array tapi ketika dilakukan inisialiasi kita
tidak menentukan size nya.

untuk basic slice kalian bisa baca dari sini ya: https://appliedgo.net/slices/#:~:text=Share%3A,while%20being%20optimized%20for%20performance

Untuk gampang nya kita pikirkan kalau slice ini berupa array yang dinamis(bukansih).
*/

func main() {
	// Maksud dari dinamis itu seperti apa ? Nah kita coba liat case array vs slice:
	/*
		   	Untuk inisialisasi slice kita bisa menggunakan dengan cara berikut
		   	var slice1 []int
		  	var slice2 = []int{}
			slice4
	*/
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// nah jika kita coba melakukan operasi seperti ini
	// array[10] = 11 <- Ini akan error
	fmt.Println(len(array), cap(array))

	// Nah kita coba liat case slice dibawah ini
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// nah jika kita coba melakukan operasi seperti ini
	// slice[10] = 11 // <- ini akan error juga karena kalau kita lihat len dan cap nya
	fmt.Println(len(slice), cap(slice)) // Output: 10 10

	// Terus gimana dong caranya ?. pada tipe slice ada method append().
	// kita bisa menambahkan data baru ke dalam slice dengan cara berikut
	slice = append(slice, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(slice) // Output: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	// Lalu kita bisa melihat len dan cap pada slice tersebut
	fmt.Println(len(slice), cap(slice)) // Output: 20 20

	// Coba coba deklrasi dengan cara berbeda
	var slice2 = []int{}
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice2)
}
