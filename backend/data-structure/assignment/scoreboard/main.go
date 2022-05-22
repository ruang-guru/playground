package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	calcScore := func(x int) int {
		return s[x].Correct*4 - s[x].Wrong
	}
	scoreI := calcScore(i)
	scoreJ := calcScore(j)

	if scoreI < scoreJ {
		return true
	} else if scoreI == scoreJ {
		if s[i].Correct < s[j].Correct {
			return true
		} else if s[i].Correct == s[j].Correct {
			if s[i].Name > s[j].Name {
				return true
			}
		}
	}
	return false // TODO: replace this
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	for i := 0; i < s.Len()-1; i++ {
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(i, j) {
				s.Swap(i, j)
			}
		}

	}

	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},  //10
		{"Agus", 3, 4, 0},  //8
		{"Doni", 3, 0, 7},  //12
		{"Ega", 3, 0, 7},   //12
		{"Anton", 2, 0, 5}, //8
	})
	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
}
