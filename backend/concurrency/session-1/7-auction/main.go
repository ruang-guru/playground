package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)            // Jumlah processor yang digunakan
	rand.Seed(time.Now().UnixNano()) // inisialisasi seed random generator
	authorized := make(chan bool)
	bidAccepted := make(chan bool)
	limit := 1000
	currentBidChan := make(chan int)
	go bid(bidAccepted, currentBidChan)
	go authorize(authorized, limit, currentBidChan)
	fmt.Println("num of goroutine before auction: ", runtime.NumGoroutine())
	if <-authorized && <-bidAccepted {
		fmt.Println("purchase success")
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("num of goroutine 10ms after auction: ", runtime.NumGoroutine())
}

func authorize(authorized chan<- bool, limit int, currentBidChan <-chan int) {
	fmt.Println("authorizing payment limit")
	if limit >= <-currentBidChan {
		fmt.Println("success authorizing")
		authorized <- true
		return
	}
	fmt.Println("rejected")
	os.Exit(3)
}

func bid(bidAccepted chan<- bool, currentBidChan chan<- int) {
	highestBid := rand.Intn(1000)
	var currentBid int
	var command string
	loop := true
	fmt.Println("highest bid : ", highestBid)
	for loop {
		fmt.Println("Enter bid offer")
		fmt.Scanln(&currentBid)
		if currentBid > highestBid {
			fmt.Println("bid accepted")
			currentBidChan <- currentBid
			bidAccepted <- true
			return
		}
		fmt.Println("bid rejected\nbid again ? (y/n)")
		fmt.Scanln(&command)
		if command != "y" {
			loop = false
		}
	}
	fmt.Println("cancelled")
	os.Exit(3)
}
