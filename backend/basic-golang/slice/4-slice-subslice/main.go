package main

import "fmt"

// Disini kita akan belajar tentang subslice
// Subslice adalah sebuah slice yang dapat dibuat dari sebuah slice/array lainnya.

// baseContainer[low : high]       // two-index form
// baseContainer[low : high : max] // three-index form
func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	s0 := a[:]      // <=> s0 := a[0:7:7]
	fmt.Println(s0) // Output: [0 1 2 3 4 5 6]
	s1 := s0[3:5]   // <=> s4 := s0[3:5:7]
	fmt.Println(s1) // Output: [3 4]
	s2 := s0[:2:2]  // <=> s5 := s0[3:5:5]
	fmt.Println(s2) // Output: [0 1]

	// Nah ketika kita menggunakan ini kita harus hati hati
	// Karena kita melakukan copy slice by refence.
	s1[0] = 100
	// Kalian bisa lihat bahwa s0 juga berubah value nya
	fmt.Println(s0, s1) // Output: [0 1 2 100 4 5 6] [100 4]

	// Untuk mengatasi ini kita bisa menggunakan function append
	// sebenarnya ada cara juga menggunakan copy(dest, src)
	// Tapi untuk simpel nya kita bisa menggunakan append

	var copySlice []int
	// Oh ia kita perlu meggunakan `...` untuk mengambil semua nilai dari slice
	copySlice = append(copySlice, s0...)
	fmt.Println(copySlice)
	copySlice[0] = 200
	fmt.Println(copySlice, s0) // Output: [200 1 2 100 4 5 6] [0 1 2 100 4 5 6]
}
