package main

import (
	"fmt"
	"time"
)

func main() {
	flag := [2]bool{false, false}
	var turn int
	counter := 0

	go func() {
		for n := 0; n < 4; n++ {
			fmt.Println("goroutine 1: ingin resource")
			flag[0] = true // goroutine 1 ingin resource
			turn = 1       // giliran goroutine 2
			fmt.Println("goroutine 1: menunggu resource")
			for flag[1] == true && turn == 1 { //tunggu jika goroutine lain juga ingin resource
			}
			//critical section
			counter++
			fmt.Println("goroutine 1: increment counter to ", counter)
			time.Sleep(2 * time.Millisecond)
			//end of critical section
			flag[0] = false //goroutine 1 sudah tidak ingin resource
			fmt.Println("goroutine 1: sudah selesai")

		}
	}()

	go func() {
		for n := 0; n < 4; n++ {
			fmt.Println("goroutine 2: ingin resource")
			flag[1] = true // goroutine 2 ingin resource
			turn = 0       // giliran goroutine 1
			fmt.Println("goroutine 2: menunggu resource")
			for flag[0] == true && turn == 0 { //tunggu jika goroutine lain juga ingin resource
			}
			//critical section
			counter++
			fmt.Println("goroutine 2: increment counter to ", counter)
			time.Sleep(2 * time.Millisecond)
			//end of critical section
			flag[1] = false //goroutine 2 sudah tidak ingin resource
			fmt.Println("goroutine 2: sudah selesai")
		}
	}()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(counter)
}
