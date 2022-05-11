package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int, 1)
	channel2 := make(chan int, 1)

	go func() {
		fmt.Println("goroutine 1 started")
		channel1 <- 1      // op1
		data := <-channel2 // op2
		fmt.Printf("receiving %d from channel 2\n", data)
	}()

	go func() {
		fmt.Println("goroutine 2 started")
		channel2 <- 1      // op3
		data := <-channel1 // op4
		fmt.Printf("receiving %d from channel 1\n", data)
	}()
	time.Sleep(100 * time.Millisecond) // agar pindah ke goroutine yang dibuat
}
