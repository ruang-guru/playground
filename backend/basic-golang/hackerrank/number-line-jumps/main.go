package main

import (
	"fmt"
)

// Number Line Jumps
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/kangaroo/problem

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	//beginanswer
	count := 0
	check := true
	if (x1 < x2) && ((v1 < v2) || (v1 == v2)) {
		return "NO"
	}
	for check {
		count++
		x1 += v1
		x2 += v2
		if x1 == x2 {
			check = false
			return "YES"
		}
		if count == 10000 {
			return "NO"
		}
	}
	return "NO"
	//endanswer
}

func main() {
	fmt.Println(kangaroo(0, 2, 5, 3))
}
