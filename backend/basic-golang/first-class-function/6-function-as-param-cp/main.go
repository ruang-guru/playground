package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
	// TODO: answer here
	printString :=func (f func ()string)  {
		r := f()
		fmt.Println(r)

	}

	goodNight:=func () string {
		return "Selamat malam"
	}



	printString(goodNight)

}
