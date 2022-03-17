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
	var i, j, k, l int
	for i = size; i > 0; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf(" ")
		}
		for k = size; k > i-1; k-- {
			fmt.Printf("*")
		}
		for l = size; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	// for h = 0; h < size*2; h++ {
	// 	fmt.Printf("*")
	// }

	for i = size - 1; i > 0; i-- {
		for j = 1; j <= size; j++ {
			fmt.Printf(" ")
		}
		for k = i; k <= i*2-1; k++ {
			fmt.Printf("*")
		}

		fmt.Printf("\n")
	}
}
