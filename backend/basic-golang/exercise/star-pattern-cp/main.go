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
	for i := size; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
