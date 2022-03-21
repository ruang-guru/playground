package main

import "fmt"

// Check Point:
// Start Pattern
// - Input: Size
// - Output: Start Pattern

// Contoh:
// - Input: Size: 10
// - Output:
//           *
//          **
//         ***
//        ****
//       *****
//      ******
//     *******
//    ********
//   *********
//  **********

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

	// TODO: answer here
	for i := 0; i < size; i++ {
		for y := 0; y < size-i-1; y++ {
			fmt.Print(" ")
		}
		for y := 0; y < i+1; y++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
