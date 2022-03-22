package main

import (
	"fmt"
	"math"
)

// Diagonal Difference
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/diagonal-difference/problem

func diagonalDifference(arr [][]int32) int32 {
	// TODO: answer here
	ltrSum := int32(0)
	rtlSum := int32(0)

	for i := 0; i < len(arr); i++ {
		ltrSum += arr[i][i]
		rtlSum += arr[i][len(arr)-1-i]
	}

	return int32(math.Abs(float64(ltrSum - rtlSum)))
}

func main() {

	var arr [][]int32
	fmt.Println(diagonalDifference(arr))
}
