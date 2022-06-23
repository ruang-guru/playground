package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
		//kenapa valuenya sama semua ?
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stop")
}
