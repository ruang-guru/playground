package main

import (
	"fmt"
)

// Apple and Orange
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/apple-and-orange/problem

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// TODO: answer here
	appleCount, orangeCount := 0, 0
	for _, val := range apples {
		dist := val + a
		if dist >= s && dist <= t {
			appleCount++
		}
	}
	for _, val := range oranges {
		dist := val + b
		if dist >= s && dist <= t {
			orangeCount++
		}
	}

	fmt.Printf("%d\n%d", appleCount, orangeCount)
}

func main() {
	var apples = []int32{3, 2, -2}
	var oranges = []int32{2, 1, 5, -6}

	countApplesAndOranges(7, 11, 5, 15, apples, oranges)
}
