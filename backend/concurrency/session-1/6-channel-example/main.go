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

func getChar(str string, c chan<- bool) {
	for _, char := range str {
		fmt.Printf("%c at time %v\n", char, time.Since(start))
	}
	c <- true
}

func getDigit(digits []int, c <-chan bool) {
	<-c
	for _, digit := range digits {
		fmt.Printf("%d at time %v\n", digit, time.Since(start))
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Main started at time ", time.Since(start))

	c := make(chan bool)
	go getChar("abcde", c)
	go getDigit([]int{1, 2, 3, 4, 5}, c)

	fmt.Println("Before sleep ", time.Since(start))
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main stopped at time ", time.Since(start))
}
