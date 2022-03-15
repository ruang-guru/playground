package main

import "fmt"

// Sama hal nya dengan array pada slice kita dapat mengubah array pada index tertentu

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice) // output: [1 2 3 4 5 6 7 8 9 10]
	slice[0] = 100
	slice[9] = 1000
	fmt.Println(slice) // output: [100 2 3 4 5 6 7 8 9 1000]
}
