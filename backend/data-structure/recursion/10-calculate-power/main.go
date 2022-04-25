// Golang program to calculate the power of a
// given number using recursion

package main

import "fmt"

func CalculatePower(num int, power int) int {
	var result int = 1
	if power > 0 {
		result = num * (CalculatePower(num, power-1))
	}
	return result
}

func main() {
	var base, power int
	var result int

	fmt.Printf("Enter value of base: ")
	fmt.Scan(&base)

	fmt.Printf("Enter value of power: ")
	fmt.Scan(&power)

	result = CalculatePower(base, power)

	fmt.Printf("%d to the power of %d is: %d\n", base, power, result)
}
