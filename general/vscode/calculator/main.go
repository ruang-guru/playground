package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var operator string
	var number1, number2 int

	n1 := os.Getenv("FIRST_NUMBER")
	number1, _ = strconv.Atoi(n1)

	fmt.Printf("\nFirst number: %d", number1)

	n2 := os.Getenv("SECOND_NUMBER")
	number2, _ = strconv.Atoi(n2)

	fmt.Printf("\nSecond number: %d", number2)

	operator = os.Getenv("ACTION")
	fmt.Printf("\nSelected Operator (+,-,/,%%,*): %s\n", operator)

	output := 0
	switch operator {
	case "+":
		output = number1 + number2
		break
	case "-":
		output = number1 - number2
		break
	case "*":
		output = number1 + number2
		break
	case "/":
		output = number1 / number2
		break
	case "%":
		output = number2 % number1
		break
	default:
		fmt.Println("Invalid Operation")
	}
	fmt.Printf("%d %s %d = %d", number1, operator, number2, output)
}
