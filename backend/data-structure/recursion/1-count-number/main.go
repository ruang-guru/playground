package main

import "fmt"

var count = 0

func recursivemethod() {
	count++
	if count <= 10 {
		fmt.Println(count)
		recursivemethod()
	}
}

func main() {
	recursivemethod()
}
