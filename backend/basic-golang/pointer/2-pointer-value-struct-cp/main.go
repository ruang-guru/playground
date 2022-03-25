package main

import (
	"fmt"
)

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// SetWidthPointer(width int) dan SetWidthValue(width int)
// SetWidthPointer(width int) untuk mengubah width dengan pointer receiver
// SetWidthValue(width int) untuk mengubah width dengan value

// TODO: answer here
type Rectangle struct {
	// TODO: answer here
	Width  int
	Length int
}

func (v *Rectangle) SetWidthPointer(width int) {
	v.Width = width
}
func (i Rectangle) SetWidthValue(width int) {
	i.Width = width
}
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.Width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.Width)
	r.SetWidthValue(70)
	fmt.Println("sesudah melakukan set width dengan value", r.Width)
}
