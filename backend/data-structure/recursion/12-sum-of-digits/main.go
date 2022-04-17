// Golang program to calculate the sum of all digits
// of a given number using recursion

package main

import "fmt"

var sum int = 0

func SumOfDigits(num int) int {
	if num > 0 {
		sum += (num % 10) //add digit into sum
		SumOfDigits(num / 10)
	}
	return sum
}

func main() {
	var num int = 0
	var result int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	result = SumOfDigits(num)

	fmt.Printf("Sum of digits is: %d\n", result)
}
