package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // menambahkan 1 ke counter
		go func(i int) {
			defer wg.Done() //mengurangi 1 dari counter
			fmt.Println(i)
		}(i)
	}
	fmt.Println("blocking")
	wg.Wait() //menunggu counter pada WaitGroup menjadi 0
}

//kadang tulisan blocking tidak muncul paling awal
//hal ini terjadi jika menggunakan multi-processor
