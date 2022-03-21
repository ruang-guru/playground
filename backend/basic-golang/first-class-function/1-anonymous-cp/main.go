package main

import "fmt"

func main() {
	//fungsi goodMorning melakukan print "selamat pagi"
	// TODO: answer here
	goodMorning := func() {
		fmt.Println("Selamat Pagi")
	}
	goodMorning()
	fmt.Printf("jenis variabelnya %T", goodMorning)
}
