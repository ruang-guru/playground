package main

import (
	"fmt"
	"time"
)

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
	// empty struct tidak perlu alokasi memori
	// channel dan empty struct biasa digunakan hanya sebagai signal penanda
	// sesuatu yang kita inginkan
	// https://golangbyexample.com/empty-struct-in-go/
}

func newSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

//request token
func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

//mengembalikan token
func (s *semaphore) Release() {
	<-s.semC
}

//kita ingin hanya ada 5 goroutine yang berjalan bersamaan
func main() {
	sem := newSemaphore(5)
	doneC := make(chan bool, 1) //hanya untuk blocking pada main goroutine
	totProcess := 10
	for i := 1; i <= totProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningProcess(v)
			if v == totProcess {
				close(doneC)
				//memberi tahu main goroutine kalau goroutine
				//lain sudah selesai
			}
		}(i)
	}
	<-doneC
}
func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(2 * time.Second) // melakukan sesuatu
}

// referensi dari https://levelup.gitconnected.com/go-concurrency-pattern-semaphore-9587d45f058d
