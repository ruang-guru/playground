package main

import (
	"fmt"
	"sync"
)

//kita bisa menggunakan wg untuk menunggu goroutine selesai berjalan
func counter(output chan<- int) {
	var wg sync.WaitGroup

	count := 0
	//angka pada comment dibawah merupakan urutan berjalannya
	wg.Add(1) // 1
	go func() {
		defer wg.Done() //3
		count++
	}()

	wg.Wait() // 2. blocking disini
	output <- count
}

func main() {
	output := make(chan int)
	go counter(output)
	fmt.Println(<-output)
}
