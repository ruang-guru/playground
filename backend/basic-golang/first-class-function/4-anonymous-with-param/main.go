package main

import "fmt"

func main() {
	greet := func(nama string) string {
		return "Halo " + nama
	}("doni")

	fmt.Println(greet)
}
