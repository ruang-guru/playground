package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3}

	res := SumAll(arr)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := SumAllCorrect(arr)
	// fmt.Println(resCorrect)
}

func SumAll(arr []int) int {
	res := 0
	for val := range arr {
		res += val
	}
	return res
}

func SumAllCorrect(arr []int) int {
	return 0 // TODO: replace this
}
