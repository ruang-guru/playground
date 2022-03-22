package main

import "fmt"

func main() {
	// mengembalikan string selamat sore dengan anonymous function
	goodAfternoon := func() string {
		//beginanswer
		return "Selamat sore"
		//endanswer
	}()

	fmt.Println(goodAfternoon)
}
