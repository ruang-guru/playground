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

	// TODO: answer here
	for i := 1; i < size; i++ {
		for j := i; j < size; j++ {
			fmt.Print(" ")
		}
		for j := 2 * i - 1; j > 0; j-- {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := size; i > 0; i-- {
		for j := size - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for j := 2 * i - 1; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
