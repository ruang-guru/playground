package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	go func() {
		fmt.Println("goroutine 1 started")
		select {
		case channel1 <- 1: // op1
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return //kalau ga return bakal deadlock di op2
		}
		data := <-channel2 // op2
		fmt.Printf("receiving %d from channel 2\n", data)
	}()

	go func() {
		fmt.Println("goroutine 2 started")
		select {
		case channel2 <- 2: // op3
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return //kalau ga return bakal deadlock di op4
		}
		data := <-channel1 // op4
		fmt.Printf("receiving %d from channel 1\n", data)

	}()
	time.Sleep(500 * time.Millisecond) // agar pindah ke goroutine yang dibuat
}
