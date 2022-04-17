package main

import (
	"fmt"
	"time"
)

func main() {

	go sum(4, 5)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stop")
}

func sum(a, b int) {
	fmt.Printf("hasil penjumlahan %d dan %d: %d\n", a, b, a+b)
}
