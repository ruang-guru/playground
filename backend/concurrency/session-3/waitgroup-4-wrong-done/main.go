package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10) // menambahkan 10 ke counter
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i) // tidak dijalankan
		}(i)
		wg.Done() //mengurangi 1 dari counter
	}
	fmt.Println("blocking")
	wg.Wait() //menunggu counter pada WaitGroup menjadi 0
}
