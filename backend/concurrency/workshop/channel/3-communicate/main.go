package main

import "fmt"

func main() {
	c := make(chan int)
	go func(a, b int) {
		fmt.Println("send to channel c")
		c <- a + b // mengirim data ke channel c
	}(4, 5)

	result := <-c // menerima data dan menyimpannya ke result
	fmt.Println("hasil penjumlahan", result)
	fmt.Println("main stop")
}
