package main

import "fmt"

// fungsi ini digunakan untuk menyapa dua nama yang ada di parameter
// kalimat sapaan yang dibuat akan dikirim melalui channel
func greet(name1, name2 string, c chan string) {
	greeting := fmt.Sprintf("halo %s dan %s", name1, name2)
	fmt.Printf("kirim %s ke channel\n", greeting)
	// TODO: answer here
}

//gunakan channel untuk komunikasi antar goroutine
func start(name1, name2 string, output chan string) {
	c := make(chan string)
	greeting := ""
	fmt.Printf("akan menyapa %s dan %s\n", name1, name2)

	// menjalankan goroutine greet
	// TODO: answer here

	//menerima data dari channel
	// TODO: answer here
	fmt.Println(c) //agar variabel c digunakan
	fmt.Printf("fungtion greet say: %s\n", greeting)
	output <- greeting
}
