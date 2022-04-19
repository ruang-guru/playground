// [Golang] Calculate Least Common Multiple (LCM) by GCD

// Find least common multiple (LCM) by greatest common divisor (GCD) via Go programming language.

// We use the following formula to calculate LCM:

// For a, b \in \mathbb{N}, a * b = LCM(a,b) * GCD(a,b)
// To calculate GCD, see my previous post [1].

// Multiples of 3
// (0), 3, 6, 9, (12), 15, 18, 21, (24), ...

// Multiples of 4
// (0), 4, 8, (12), 16, 20, (24), 28, 32, ...

// The LCM of 3 and 4 is 12.

package main

import (
	"fmt"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	// if there are more integers, find LCM of all integers
	for i := 0; i < len(integers); i++ {
		// TODO: answer here
	}

	return result
}

func main() {
	fmt.Println(LCM(3, 4))
	fmt.Println(LCM(10, 15, 20))
	fmt.Println(LCM(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
