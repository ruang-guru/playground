package main

import "fmt"

func main() {
	/*
		Calculate average score

		Sample Input/Output
		[90,90,100] -> 93.333336
		[100,100,100,100] -> 100
		[] -> 0
	*/
	arr := []int{90, 90, 100}
	res := getAverageScore(arr)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := getAverageScoreCorrect(arr)
	// fmt.Println(resCorrect)
}

func getAverageScore(score []int) float32 {
	sum := 0
	for i := 0; i < len(score); i++ {
		sum += score[i]
	}
	return float32(sum) / float32(len(score))
}

func getAverageScoreCorrect(score []int) float32 {
	if len(score) == 0 {
		return 0
	}
	sum := 0

	for i := 0; i < len(score); i++ {
		sum += score[i]
	}
	return float32(sum) / float32(len(score))
}
