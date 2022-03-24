package main

import "fmt"

// buat struct Rectangle (persegi panjang) dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan (print) luas persegi panjang
// GetPerimeter() digunakan untuk menampilkan (print) keliling persegi panjang

type Rectangle struct {
	// TODO: answer here
	Width  int
	Length int
}

func (r Rectangle) GetArea(luas int) {
	fmt.Println(luas)
}
func (r Rectangle) GetPerimeter(keliling int) {
	fmt.Println(keliling)
}

// TODO: answer here
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea(r.Length * r.Width)
	r.GetPerimeter(2*r.Length + 2*r.Width)
}
