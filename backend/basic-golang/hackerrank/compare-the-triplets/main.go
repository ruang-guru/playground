package main

import (
	"fmt"
)

// Compare the Triplets
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/compare-the-triplets/problem

func compareTriplets(a []int32, b []int32) []int32 {

	// TODO: answer here
	score := []int32{0, 0}

	for index, val := range a {
		if val > b[index] {
			score[0]++
		} else if val < b[index] {
			score[1]++
		}
	}
	return score
}

func main() {

	var a = []int32{5, 6, 7}
	var b = []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b))
}
