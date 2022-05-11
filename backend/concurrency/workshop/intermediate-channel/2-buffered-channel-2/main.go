package main

import (
	"fmt"
	"time"
)

//mengembalikan hasil pangkat dua angka 1-10
//dapat dilihat tidak terjadi blocking selama buffer belum penuh
func squareWorker(output chan<- int) {
	for i := 1; i < 11; i++ {
		output <- i * i
	}
	fmt.Println("selesai mengirim")
}

func main() {
	output := make(chan int, 10)
	go squareWorker(output)
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10; i++ {
		fmt.Println("main receive output:", <-output)
	}
}
