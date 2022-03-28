package main

import "fmt"

func main() {
	// mengembalikan string selamat sore dengan anonymous function
	goodAfternoon := func() string {
		// TODO: answer here
		return "selamat sore"
	}()

	fmt.Println(goodAfternoon)
}
