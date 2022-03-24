package main

import "fmt"

// buat struct Rectangle (persegi panjang) dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan (print) luas persegi panjang
// GetPerimeter() digunakan untuk menampilkan (print) keliling persegi panjang

type Rectangle struct {
	// TODO: answer here
}

// TODO: answer here
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea()
	r.GetPerimeter()
}
