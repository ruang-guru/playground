package main

import "fmt"

// Jadi contoh ini kita akan menggunakan composite pattern disalah satu problem klasik
// Yaitu sistem Folder dan File
// Nah anggak saja Folder itu adalah sebuah `Kardus` dan `File` itu sebuah `Produk`
// Tetapi kasusnya disini kita mencari sebuah keyword di dalamnya

type File struct {
	Name string
	Byte string // Anggaplah byte ini adalah isi dari sebuah file tersebut
}

// Sama seperti di folder, File juga memiliki interface `Component`
// Pada file ini kita tidak menggunakan for loop karena kita tahu bahwa `File` merupakan unit terkecil dari sebuah `Component`
func (f *File) Search(keyword string) bool {
	fmt.Printf("Searching for keyword %s in File: %s\n", keyword, f.Name)
	if f.Byte == keyword {
		fmt.Printf("Found keyword %s in File: %s\n", keyword, f.Name)
		return true
	}
	return false
}

func (f *File) getName() string {
	return f.Name
}
