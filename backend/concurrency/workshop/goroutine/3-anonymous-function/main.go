package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello from goroutine")
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stop")
}
