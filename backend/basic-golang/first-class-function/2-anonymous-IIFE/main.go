package main

import "fmt"

func main() {
	// fungsi yang lansung dijalankan
	func() {
		fmt.Println("Halo dari anonymous function")
	}()

	//contoh lain untuk melakukan pemangkatan 2
	func(x int) {
		fmt.Println(x * x)
	}(3)

}
