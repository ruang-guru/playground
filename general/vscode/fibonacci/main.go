package main

import (
	"fmt"
	"os"
	"strconv"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return 0
	}
	prev := 0
	curr := 1
	for i := 0; i < n-2; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func main() {
	n := os.Getenv("NUMBER")
	number, _ := strconv.Atoi(n) // convert string to integer
	result := Fibonacci(number)
	fmt.Printf("\nFibonacci of %d is %d\n", number, result)

}
