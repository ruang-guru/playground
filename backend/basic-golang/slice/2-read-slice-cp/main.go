package main

import "fmt"

// Disini kalian coba untuk membaca nilai dari slice yang diberikan.
// Lalu kalian akan menambahkan kata "Olleh" pada slice tersebut.
func main() {
	slice := []string{"Hello", "World"}
	slice = append(slice, "Olleh")
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice)
}
