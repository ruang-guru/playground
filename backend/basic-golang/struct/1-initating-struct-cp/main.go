package main

import "fmt"

//buat struct Rectangle dengan dua atribut yaitu Width dan Length

type Rectangle struct {
	Width  int64
	Length int64
}

func (p *Rectangle) SetWidth(width int64, length int64) {
	p.Width = width
	p.Length = length
}

func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r) //{10,20}

	r2 := Rectangle{Width: 5, Length: 15}
	fmt.Println(r2) // {5,15}

	fmt.Println("lebar rectangle2:", r2.Width)    //5
	fmt.Println("panjang rectangle2:", r2.Length) //15

}
