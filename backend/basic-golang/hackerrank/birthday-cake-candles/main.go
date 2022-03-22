package main

import (
	"fmt"
)

// Birthday Cake Candles
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func birthdayCakeCandles(candles []int32) int32 {
	// TODO: answer here
	max, count := int32(-1), int32(0)
	for _, val := range candles {
		if val > max {
			max = val
			count = int32(0)
		}
		if val == max {
			count++
		}
	}
	return count
}

func main() {
	var arr = []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(arr))
}
