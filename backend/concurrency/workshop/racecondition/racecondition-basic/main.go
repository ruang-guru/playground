package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			counter += 1
		}()

		go func() {
			counter += 2
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(counter)
}
