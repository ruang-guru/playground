package main

import "fmt"

// Demo:
// Start Pattern
// - Input: Size: 10
// - Output:
// *
// **
// ***
// ****
// *****
// ******
// *******
// ********
// *********
// **********

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

	for i := 0; i < size; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
}
