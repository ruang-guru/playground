package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat bereksperimen dengan zero value pada pointer.
// Program dibawah akan mengeluarkan panic, modifikasi kode dibawah sehingga program tidak mengeluarkan panic.

type Student struct {
	Name         string
	Class        string
	SubjectScore *SubjectScore
}

type SubjectScore struct {
	Scores []int
}

func GetMaxScore(student *Student) int {
	// TODO: answer here
	maxScore := 0
	for _, score := range student.SubjectScore.Scores {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	student := Student{
		Name:  "Tony",
		Class: "X IPA",
	}

	maxScore := GetMaxScore(&student)
	fmt.Println(maxScore)
}
