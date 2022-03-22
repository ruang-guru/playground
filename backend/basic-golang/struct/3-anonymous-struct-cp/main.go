package main

import "fmt"

func main() {
	//membuat rectangle dengan anonymous struct
	//field dari struct ini sama seperti rectangle sebelumnya
	//beginanswer
	newRectangle := struct {
		Width  int
		Length int
	}{
		Width:  10,
		Length: 30,
	}
	//endanswer

	fmt.Printf("rectangle dengan lebar %d dan panjang %d", newRectangle.Width, newRectangle.Length)
}
