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
		for j := 0; j <size-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
		
	}

}
