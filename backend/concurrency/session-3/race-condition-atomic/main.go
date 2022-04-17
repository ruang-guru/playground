package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var a int32 = 0
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&a, 1)
		}()
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(a) // apa nilai dari variabel a?
}
