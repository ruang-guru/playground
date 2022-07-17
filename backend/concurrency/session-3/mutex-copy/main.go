package main

import (
	"fmt"
	"sync"
	"time"
)

type count struct {
	mu sync.Mutex
}

func (c count) lock() {
	fmt.Printf("mutex address: %p\n", &c.mu)
	c.mu.Lock()
}

func (c count) unlock() {
	fmt.Printf("mutex address: %p\n", &c.mu)
	c.mu.Unlock()
}

func main() {
	counter := 0
	a := count{}
	for i := 0; i < 3; i++ {
		go func() {
			a.lock()
			counter++
			a.unlock() // error disini
			// karena default state mutex adalah unlocked
			// saat melakukan unlock pada unlocked mutex akan terjadi error
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println(counter)
}

//referensi https://eli.thegreenplace.net/2018/beware-of-copying-mutexes-in-go/
