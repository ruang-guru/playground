package main

import "fmt"

func main() {
	order("teh")
	order("kopi")
}

func order(drink string) {
	// mengucapkan terima kasih di akhir dengan menggunakan defer
	// TODO: answer here
	defer fmt.Print("terima kasih \n\n")
	fmt.Println("kami sedang ada diskon untuk pembelian kopi")
	fmt.Println("pesanan anda:", drink)
	if drink == "kopi" {
		fmt.Println("okee, karena diskon totalnya jadi 4000")
	}
	fmt.Println("mohon ditunggu")
}
