package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
	// TODO: answer here

	getAreaFunc := func() func(x, y int) int {
		return func(x, y int) int {
			return x * y
		}
	}

	areaF := getAreaFunc()
	res := areaF(3, 4) // 12
	fmt.Println(res)
}
