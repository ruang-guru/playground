package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

//Pada program ini kita melakukan pemanggilan API secara sequential
//Bagaimana cara agar program ini dapat berjalan secara concurrent ?
func main() {
	APICallA()
	APICallB()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("from main function at time", time.Since(start))
}

func APICallA() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("from APICallA at time", time.Since(start))
}

func APICallB() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("from APICallB at time", time.Since(start))
}
