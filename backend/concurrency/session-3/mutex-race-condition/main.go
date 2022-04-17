package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := &sync.Mutex{}
	a := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(a) // apa nilai dari variabel a?
}
