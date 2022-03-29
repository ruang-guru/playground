package main

import "fmt"

// Demo:
// Pyramid Pattern
// - Input: Size: 10
// - Output:
//           *
//          ***
//         *****
//        *******
//       *********
//      ***********
//     *************
//    ***************
//   *****************
//  *******************

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

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
}
