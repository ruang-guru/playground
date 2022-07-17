package main

import (
	"fmt"
)

//goroutine dapat berjalan secara concurrent
//disini goroutine akan mengcapture variable i secara reference
//ini perlu dilakukan agar goroutine dapat mengubah value i
//bisa disebut closure juga karena fungsi goroutine reference variable di luar scopenya
func main() {
	i := 0 //nilai i akan diubah oleh goroutine
	go func() {
		for i < 10 {
			fmt.Println("increment", i)
			i++
		}
	}()

	for i < 10 {
		fmt.Println("loop")
	}
	fmt.Println("main stop")
}
