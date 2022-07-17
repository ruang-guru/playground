package main

import "fmt"

//buferred channel tidak blocking jika buffer belum penuh
func main() {
	input := make(chan string, 3)

	input <- "halo"
	input <- "dari buffered"
	input <- "channel"

	for i := 0; i < cap(input); i++ {
		fmt.Println(<-input)
	}
}
