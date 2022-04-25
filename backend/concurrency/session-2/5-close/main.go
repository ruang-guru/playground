package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 5)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- "every 500 millisecond"
			time.Sleep(500 * time.Millisecond)
		}
		close(c1)
	}()

	for msg := range c1 {
		fmt.Println(msg)
	}
}
