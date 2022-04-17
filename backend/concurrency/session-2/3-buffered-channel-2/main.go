package main

import (
	"fmt"
)

func sender(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine send:", i)
		c <- i
	}
}
func main() {
	c := make(chan int, 1)
	go sender(c)
	for i := 0; i < 4; i++ {
		fmt.Println("main receive:", <-c)
	}
}
