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

// TODO: answer here
func (r Rectangle) GetArea() {
	fmt.Println("Luas : ", r.Width*r.Length)
}

func (r Rectangle) GetPerimeter() {
	fmt.Println("keliling : ", 2*r.Width+r.Length)
}

func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea()
	r.GetPerimeter()
}
