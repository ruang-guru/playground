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

func getChar(str string) {
	for _, char := range str {
		fmt.Printf("%c ", char) //view withouth time

		// fmt.Printf("%c %v\n", char, time.Since(start)) //view with time
	}
	fmt.Println()
}

func getDigits(digits []int) {
	for _, digit := range digits {
		fmt.Printf("%v ", digit) //run withouth time

		// fmt.Printf("%d %v\n", digit, time.Since(start)) //view with time
	}
	fmt.Println()
}

func main() {
	runtime.GOMAXPROCS(1) //agar berjalan parallel matikan perintah ini
	fmt.Println("Main started", time.Since(start))

	go getChar("abcde")

	go getDigits([]int{1, 2, 3, 4, 5})
	fmt.Println("Before sleep", time.Since(start))
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main stopped ", time.Since(start))
}
