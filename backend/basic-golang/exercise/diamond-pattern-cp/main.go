package main

import "fmt"

// Check Point:
// Diamond Pattern
// - Input: Size: 5
// - Output:
//         *
//        ***
//       *****
//      *******
//     *********
//    ***********
//     *********
//      *******
//       *****
//        ***
//         *

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

	//TODO
	var i, j, k int
	for i = 0; i <= size+1; i++ {
		for j = size + 1; j > i; j-- {
			fmt.Print(" ")
		}
		for k = 0; k < j; k++ {
			fmt.Print("*")
		}
		for k = 2; k < i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i = 0; i <= size; i++ {
		for j = 1; j < i+2; j++ {
			fmt.Print(" ")
		}
		for k = size; k > i; k-- {
			fmt.Print("*")
		}
		for k = size - 1; k > i; k-- {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
