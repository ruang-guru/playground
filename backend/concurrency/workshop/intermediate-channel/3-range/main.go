package main

import (
	"fmt"
	"time"
)

func main() {
	input := make(chan string, 10)

	go receiver(input)
	for i := 0; i < 10; i++ {
		input <- "every 500 millisecond"
		time.Sleep(500 * time.Millisecond)
	}

}

func receiver(input chan string) {
	for msg := range input {
		fmt.Println(msg)
	}
}

// saat data yang dikirim main goroutine habis, tidak terjadi deadlock
// ini terjadi karena main goroutine masih terus berjalan
// yang terkena block hanya goroutine yang dibuat
// ketika main goroutine selesai berjalan, maka goroutine lain juga akan dihentikan
