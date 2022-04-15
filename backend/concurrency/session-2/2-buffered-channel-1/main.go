package main

import "fmt"

func main() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	//bisa jadi contoh lanjutan untuk select/range

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
}
