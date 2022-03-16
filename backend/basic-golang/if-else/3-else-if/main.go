package main

import "fmt"

// Disini kita akan belajar tentang else dan else if

func main() {
	nilai := 100
	if nilai < 200 {
		fmt.Println("nilai lebih kecil dari 200")
	} else if nilai == 150 {
		fmt.Println("nilai sama dengan 150")
	} else if nilai == 100 {
		fmt.Println("nilai sama dengan 100")
	} else if nilai == 50 {
		fmt.Println("nilai sama dengan 50")
	} else {
		fmt.Println("ndak tau kok tanya saya")
	}
}
