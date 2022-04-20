package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex
	counterGreedy := 0
	counterPolite := 0

	// goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				counterGreedy++
				time.Sleep(100 * time.Microsecond)
				//do something to shared resource
				mu.Unlock()
			}
		}
	}()

	// goroutine 2
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		counterPolite++
		//do something to shared resource
		mu.Unlock()
	}
	done <- true
	fmt.Println("counter greedy:", counterGreedy)
	fmt.Println("counter polite:", counterPolite)

}

//goroutine1 mendapatkan lebih banyak akses daripada goroutine2
//walaupun keduanya sama" sleep selama 100 microsecond

//sumber : https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50
