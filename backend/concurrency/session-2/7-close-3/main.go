package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string, 5)

	for i := 0; i < 5; i++ {

		c1 <- fmt.Sprintf("data %d", i) //untuk membuat string dengan format
	}

	close(c1)
	close(c1)

	for msg := range c1 {
		fmt.Println(msg)
	}
}
