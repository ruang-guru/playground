package main

import "fmt"

func main() {
	/*
		Return the maximum / the biggest number, if empty return 0

		Example Input/Output
		[7,1,2,3,4,5,10] -> 10
		[] -> 0
		[-1,-3,-5,-1,-2,-3,-4,-10,-15,-5] -> -1
	*/
	arr := []int{-1, -3, -5, -1, -2, -3, -4, -10, -15, -5}
	res := findMax(arr)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := findMaxCorrect(arr)
	// fmt.Println(resCorrect)
}

func findMax(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func findMaxCorrect(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
