package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
	// TODO: answer here
	getAreaFunc := func() func(num1, num2 int) int {
		return func(num1, num2 int) int {
			return num1 * num2
		}
	}
	areaF := getAreaFunc()
	res := areaF(3, 4)  // 12
	res2 := areaF(5, 5) // 12
	fmt.Println(res)
	fmt.Println(res2)
}
