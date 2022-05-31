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
	space := size - 1

	for j := 1; j <= size; j++ {
		for i := 1; i <= space; i++ {
			fmt.Print(" ")
		}
		space--
		for i := 1; i <= 2*j-1; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	space = 1
	for j := 1; j <= size-1; j++ {
		for i := 1; i <= space; i++ {
			fmt.Print(" ")
		}
		space++
		for i := 1; i <= 2*(size-j)-1; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
