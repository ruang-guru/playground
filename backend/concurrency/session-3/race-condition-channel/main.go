package main

import (
	"fmt"
)

func main() {
	a := 0
	input := make(chan int)
	for i := 0; i < 1000; i++ {
		go func() {
			input <- 1
		}()

	}
	for i := 0; i < 1000; i++ {
		a += <-input
	}
	fmt.Println(a) // apa nilai dari variabel a?
}
