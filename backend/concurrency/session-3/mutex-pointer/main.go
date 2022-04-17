package main

import (
	"fmt"
	"sync"
	"time"
)

type count struct {
	mu *sync.Mutex
	//bisa juga langsung sync.Mutex
	//kenapa bisa begitu ? karena methodnya menggunakan pointer
}

func (c *count) lock() {
	fmt.Printf("mutex address: %p\n", &c.mu)
	c.mu.Lock()
}

func (c *count) unlock() {
	fmt.Printf("mutex address: %p\n", &c.mu)
	c.mu.Unlock()
}

func main() {
	counter := 0
	a := count{mu: &sync.Mutex{}}
	for i := 0; i < 10; i++ {
		go func() {
			a.lock()
			counter++
			a.unlock()
			//menggunakan mutex yang sama
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println(counter)
}

//referensi https://eli.thegreenplace.net/2018/beware-of-copying-mutexes-in-go/
