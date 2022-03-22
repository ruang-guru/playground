package main

import (
	"fmt"
)

// Compare the Triplets
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/compare-the-triplets/problem

func compareTriplets(a []int32, b []int32) []int32 {

	//beginanswer
	arr := []int32{0, 0}
	for i := range a {
		if a[i] > b[i] {
			arr[0] = arr[0] + 1
		} else if a[i] < b[i] {
			arr[1] = arr[1] + 1
		}
	}
	return arr
	//endanswer
}

func main() {

	var a = []int32{5, 6, 7}
	var b = []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b))
}
