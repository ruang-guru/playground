package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan blank return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

// TODO: answer here
func square(num1, num2 int) (res1, res2 int) {
	res1 = num1 * num1
	res2 = num2 * num2
	return
	
}
