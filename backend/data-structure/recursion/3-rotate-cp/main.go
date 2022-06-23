// Expected Rosated:
// [(1) 2 3 4 5 6 7 8 9 (10)] => [(10) 2 3 4 5 6 7 8 9 (1)]

// Proses Rotate:
// [10 9 3 4 5 6 7 8 2 1]
// [10 9 8 4 5 6 7 3 2 1]
// [10 9 8 7 5 6 4 3 2 1]
// [10 9 8 7 6 5 4 3 2 1]
// [10 9 8 7 5 6 4 3 2 1]
// [10 9 8 4 5 6 7 3 2 1]
// [10 9 3 4 5 6 7 8 2 1]
// [10 2 3 4 5 6 7 8 9 1]

package main

import "fmt"

func Rotate(index int, args []int) []int {
	if len(args) == 0 {
		return []int{}
	}

	if index == -1 {
		return args
	}

	if index == 0 {
		// TODO: answer here
		return Rotate(1, args)
	}

	if index == len(args)-1 {
		return Rotate(-1, args)
	}

	return Rotate(index+1, args)

}

func main() {
	mylist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rotated := Rotate(0, mylist)
	fmt.Println(rotated)
}
