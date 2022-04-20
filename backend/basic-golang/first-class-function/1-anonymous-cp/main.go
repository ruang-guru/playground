package main

import "fmt"

func main() {
	//fungsi goodMorning melakukan print "selamat pagi"
	// TODO: answer here
	goodMorning := func() {
		fmt.Println("Selamat pagi")
	}
	goodMorning()
	fmt.Printf("jenis variabelnya %T", goodMorning)
}

// func main() {
// 	greet := func() {
// 		fmt.Println("Halo dari anonymous function")
// 	}
// 	greet()
// 	fmt.Printf("jenis variabelnya %T", greet)
// }
