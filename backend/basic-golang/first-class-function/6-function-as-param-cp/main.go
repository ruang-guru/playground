package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
	// TODO: answer here
	printString := func(f func() string) {
		greet := f()
		fmt.Println(greet)
	}

	goodNight := func() string {
		return "selamat malam"
	}

	printString(goodNight)

}
