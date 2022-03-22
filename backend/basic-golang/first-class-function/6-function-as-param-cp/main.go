package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
	//beginanswer
	printString := func(f func() string) {
		result := f()
		fmt.Println(result)
	}

	goodNight := func() string {
		return "good night friends"
	}
	//endanswer

	printString(goodNight)

}
