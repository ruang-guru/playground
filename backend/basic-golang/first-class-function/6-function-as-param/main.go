package main

import "fmt"

func main() {
	printString := func(f func() string) {
		result := f()
		fmt.Println(result)
	}

	hello := func() string {
		return "hello friends"
	}

	printString(hello)

}
