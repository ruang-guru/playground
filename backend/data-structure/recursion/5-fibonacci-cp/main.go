// Fibonacci

// Problem:
// Buat fungsi untuk menghasilkan angka Fibonacci ke-n.
// The nth Fibonacci number is given by:
// F(n) = F(n-1) + F(n-2)
// F(0) = 0
// F(1) = 1
// F(2) = 1
// F(3) = 2
// F(4) = 3
// F(5) = 5
// F(6) = 8
// F(7) = 13
// F(8) = 21
// F(9) = 34
// ...

// Dua suku pertama barisan tersebut adalah 0, 1.
// Contoh: fib(0) = 0, fib(1) = 1, fib(2) = 1

package main

import "fmt"

func FibonacciRecursion(i int) int {

	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	// TODO: answer here
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf(" %d  ", FibonacciRecursion(i))
	}
}
