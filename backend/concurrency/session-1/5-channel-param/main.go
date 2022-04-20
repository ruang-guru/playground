package main

import (
	"fmt"
	"runtime"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("main started at time ", time.Since(start))
	c := make(chan string)
	go sendHello(c)

	fmt.Printf("goroutine sent this: %v. At time %v\n", <-c, time.Since(start))
	fmt.Printf("main stopped at time %v\n", time.Since(start))
}

func sendHello(c chan string) {
	c <- "hello from goroutine"
}
