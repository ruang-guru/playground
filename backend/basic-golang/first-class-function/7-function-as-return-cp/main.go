package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
<<<<<<< HEAD
	// TODO: answer here
	getAreaFunc := func() func(num1, num2 int) int {
		return func(num1, num2 int) int {
			return num1 * num2
		}
	}
=======
	//beginanswer
	getAreaFunc := func() func(int, int) int {
		return func(x, y int) int {
			return x * y
		}
	}
	//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	areaF := getAreaFunc()
	res := areaF(3, 4)  // 12
	res2 := areaF(5, 5) // 12
	fmt.Println(res)
	fmt.Println(res2)
}
