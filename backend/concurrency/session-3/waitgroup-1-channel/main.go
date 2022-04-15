package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer func() { done <- true }()
			fmt.Println(i)
		}(i)
	}
	fmt.Println("blocking")
	for i := 0; i < 10; i++ {
		<-done
	}
}
