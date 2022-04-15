package main

import (
	"fmt"
	"time"
)

//saat menggunakan channel sebagai parameter function
//kita bisa menentukan apakah kita ingin mengirim atau menerima
func main() {
	c := make(chan string)
	go receive(c)
	go send(c)
	time.Sleep(100 * time.Millisecond)
}

// <-chan disini berarti channel c hanya bisa menerima data
func receive(c <-chan string) {
	fmt.Println(<-c)
	// c <- "hello"
}

// chan<- disini berarti channel c hanya bisa mengirim data
func send(c chan<- string) {
	c <- "hello"
	// <-c
}
