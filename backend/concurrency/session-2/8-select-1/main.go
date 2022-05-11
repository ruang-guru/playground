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
	c2 := make(chan string)
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			c1 <- "every 500 millisecond"
		}
	}()

	go func() {
		for {
			time.Sleep(2000 * time.Millisecond)
			c2 <- "every 2 second"
		}

	}()
	for {
		fmt.Printf("%s channel 1. At time %s\n", <-c1, time.Since(start))
		fmt.Printf("%s channel 2. At time %s\n", <-c2, time.Since(start))
	}

}
