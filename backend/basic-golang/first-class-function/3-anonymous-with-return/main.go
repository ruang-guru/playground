package main

import "fmt"

func main() {
	greet := func() string {
		return "Halo dari anonymous function"
	}()

	fmt.Println(greet)
}
