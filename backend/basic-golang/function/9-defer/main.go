package main

import "fmt"

func main() {
	greet("citra")
	greet("cahya")
}

func greet(name string) {
	//dijalankan diakhir, sebelum fungsi selesai atau terjadi error
	defer fmt.Println("Sampai jumpa lagi", name)
	fmt.Println("Halo", name)
}
