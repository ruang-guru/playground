package main

import "fmt"

func main() {
	// print selamat pagi menggunakan anonymous function
	// TODO: answer here
	goodMorning := func() {
		fmt.Println("Good morning from anonymous function")
	}
	goodMorning()
}
