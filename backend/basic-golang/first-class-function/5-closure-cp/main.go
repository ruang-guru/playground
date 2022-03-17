package main

import "fmt"

func main() {
	//fungsi counter akan menerima (x int) dan mengembalikan fungsi
	//fungsi yang dikembalikan akan melakukan decrement nilai parameter x ketika dipanggil dan
	//mengembalikan nilai parameter x

	counter := func(x int) func() int {
		// TODO: answer here
	}
	decrement := counter(5)
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}
