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
			defer wg.Done() //mengurangi 1 dari counter
			fmt.Println(i)
		}(i)
	}
	fmt.Println("blocking")
	wg.Wait() //menunggu counter pada WaitGroup menjadi 0
}

// lebih rawan bug
// karena kita perlu selalu mengubah jumlah yang di-pass kedalam param add
// sesuai dengan jumlah goroutine yang ingin dijalankan
