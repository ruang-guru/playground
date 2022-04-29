package main

import (
	"fmt"
	"github.com/pkg/profile"
	"testing"
)

func add(a int, b int) int {
	return a + b
}

// Make a simple function to use as a profiling.
func TestFlatCum(t *testing.T) {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	data := [][2]int{{100, 106}, {200, 206}, {300, 308}, {159, 259}}

	var resultAdd int
	for i := 0; i < len(data); i++ {
		resultAdd += add(data[i][0], data[i][1])
	}

	var resultMinus int
	for i := 0; i < len(data); i++ {
		data := data[i][0] - data[i][1]
		resultMinus += data
	}

	fmt.Println(resultAdd, resultMinus)
}
