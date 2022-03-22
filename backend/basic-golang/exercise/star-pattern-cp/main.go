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

<<<<<<< HEAD
	// TODO: answer here
	for i := 0; i < size; i++ {

		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < i+1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
=======
	//beginanswer
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
	//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
}
