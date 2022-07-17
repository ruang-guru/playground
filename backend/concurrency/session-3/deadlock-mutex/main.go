package main

import (
	"sync"
	"time"
)

type data struct {
	sharedData int
	mu         *sync.Mutex
}

//agar errornya keliatan, akan digunakan 1 spawned goroutine dan main goroutine
//jika dijalankan pada dua spawned goroutine
//deadlock tidak kelihatan karena main goroutine tetap berjalan
func main() {
	sharedVariable1 := &data{sharedData: 5, mu: &sync.Mutex{}}
	sharedVariable2 := &data{sharedData: 5, mu: &sync.Mutex{}}

	go func() {
		sharedVariable2.mu.Lock() // op5
		sharedVariable2.sharedData++
		sharedVariable1.mu.Lock() // op6
		sharedVariable1.sharedData++
		sharedVariable1.mu.Unlock() // op7
		// do something
		sharedVariable2.mu.Unlock() // op8
	}()

	sharedVariable1.mu.Lock() // op1
	sharedVariable1.sharedData++
	time.Sleep(100 * time.Millisecond) // agar pindah ke goroutine yang dibuat
	sharedVariable2.mu.Lock()          // op2
	sharedVariable2.sharedData++
	sharedVariable2.mu.Unlock() // op3
	// do something
	sharedVariable1.mu.Unlock() // op4
}
