package main

// Pada umumnya stack memiliki dua propertii yaitu top dan data.
// Top merupakan atribut yang menunjukkan indeks elemen teratas.
// Data merupakan atribut yang menyimpan sekumpulan data yang disimpan.
// Ketika stack tidak memiliki elemen, maka nilai top dapat diassign ke -1.

type Stack struct {
	Top  int
	Data []int
}
