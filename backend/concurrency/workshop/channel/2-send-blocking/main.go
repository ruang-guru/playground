package main

import "fmt"

//karena sifat channel berukuran 0 yang blocking
//dapat digunakan channel untuk menunggu goroutine
func main() {
	c := make(chan string)
	go func() {
		fmt.Println("receive from main")
		fmt.Println(<-c) // blocking hingga ada yang mengirim data
	}()

	c <- "main say hello" //blocking hingga data yang dikirim ada yang menerima
	// karena terjadi blocking, goroutine #2 akan dijalankan
	// ketika goroutine #2 menerima dari channel c, maka main goroutine jalan kembali
	fmt.Println("main stop")
}
