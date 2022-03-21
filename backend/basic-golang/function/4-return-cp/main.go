package main

import "fmt"

//pakai * untuk melakukan perkalian
//misal number1 * number2
func main() {
	result := multiply(3, 3)
	fmt.Println(result)
	fmt.Println(multiply(5, 5))
}

// TODO: answer here
func multiply(num1, num2 int) (res int) {
	res = num1 * num2
	return
}
