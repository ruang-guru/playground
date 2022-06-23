package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	c1 := make(chan string)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			c1 <- "every 500 millisecond"
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("%s channel 1. At time %s\n", msg, time.Since(start))
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("default here. At time", time.Since(start))
		}
	}
}
