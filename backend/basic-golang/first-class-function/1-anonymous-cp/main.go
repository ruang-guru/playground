package main

import "fmt"

func main() {
	//fungsi goodMorning melakukan print "selamat pagi"
	//beginanswer
	goodMorning := func() {
		fmt.Println("selamat pagi")
	}
	//endanswer
	goodMorning()
	fmt.Printf("jenis variabelnya %T", goodMorning)
}
