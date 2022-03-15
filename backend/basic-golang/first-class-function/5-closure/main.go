package main

import "fmt"

func main() {
	square := func(number int) func() int {
		return func() int {
			number *= number
			return number
		}
	}
	twoSquared := square(2)
	fmt.Println(twoSquared())
	fmt.Println(twoSquared())
	fmt.Println(twoSquared())
}
