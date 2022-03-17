package main

import "fmt"

func main() {
	greet := func() {
		fmt.Println("Halo dari anonymous function")
	}
	greet()
	fmt.Printf("jenis variabelnya %T", greet)
}
