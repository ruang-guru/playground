package main

import "fmt"

//karena sifat channel berukuran 0 yang blocking
//dapat digunakan channel untuk menunggu goroutine
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("send to channel c")
		c <- true // blocking hingga ada yang menerima data
	}()

	<-c //blocking
	// karena terjadi blocking, goroutine #2 akan dijalankan
	// ketika goroutine #2 mengirim ke channel c, maka main goroutine jalan kembali
	fmt.Println("main stop")
}
