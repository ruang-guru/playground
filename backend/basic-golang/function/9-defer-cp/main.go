package main

import "fmt"

func main() {
	order("teh")
	order("kopi")
}

func order(drink string) {
	// mengucapkan terima kasih di akhir dengan menggunakan defer
<<<<<<< HEAD
	// TODO: answer here
	defer fmt.Printf("Terima kasih telah memesan %s\n", drink)
	fmt.Println("\nkami sedang ada diskon untuk pembelian kopi")
=======
	//beginanswer
	defer fmt.Println("terima kasih, silahkan datang kembali")
	//endanswer
	fmt.Println("kami sedang ada diskon untuk pembelian kopi")
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	fmt.Println("pesanan anda:", drink)
	if drink == "kopi" {
		fmt.Println("okee, karena diskon totalnya jadi 4000")
	}
	fmt.Printf("mohon ditunggu\n\n")
}
