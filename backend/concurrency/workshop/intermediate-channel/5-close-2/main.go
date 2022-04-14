package main

import (
	"fmt"
	"time"
)

//Data dalam buffered channel masih bisa diambil
//walaupun channel sudah di close
func main() {
	output := make(chan string, 10)

	go sender(output)
	time.Sleep(100 * time.Millisecond)
	for msg := range output {
		fmt.Println(msg)
	}
}

func sender(output chan string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("sending %d to channel output\n", i)
		output <- fmt.Sprintf("buffer number %d", i)
	}
	close(output)
	fmt.Println("channel closed")
}
