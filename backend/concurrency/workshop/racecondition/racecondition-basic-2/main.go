package main

import (
	"fmt"
	"time"
)

//kita bisa mengatasi race condition dengan menggunakan channel
//dengan ini nilai yang didapat akan selalu 3000
func main() {
	c := make(chan int)
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			c <- 1
		}()

		go func() {
			c <- 2
		}()
	}
	go func() {
		for {
			counter += <-c
		}
	}()
	time.Sleep(200 * time.Millisecond)
	fmt.Println(counter)
}
