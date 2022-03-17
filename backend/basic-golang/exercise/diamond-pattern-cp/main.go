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
	var size,n int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)


	// TODO: answer here

	for i := 1; i <= size; i++ {
		n = 0
		for space := 1; space <= size-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			n++
			if n == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
	var i, j int
	for i = size; i >= 1; i-- {
		for space := 1; space <= size-i; space++ {
			fmt.Print("  ")
		}
		for j = i; j <= 2*i-1; j++ {
			fmt.Printf("* ")
		}
		for j = 0; j < i-1; j++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}

}
