package main

import "fmt"

// Demo:
// - Simple Pattern
// - Input: Size: 10
// - Output: **********

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

	for i := 0; i < size; i++ {
		fmt.Printf("%s", "*")
	}
	fmt.Println()
}
