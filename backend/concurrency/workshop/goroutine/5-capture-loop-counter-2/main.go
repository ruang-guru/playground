package main

import (
	"fmt"
	"time"
)

//value percobaan sebelumnya sama semua karena
//goroutine menggunakan closure yang sifatnya reference
//jadi ketika goroutine berjalan, nilai variabel i sudah berubah menjadi 5
//lalu bagaimana caranya agar i yang dicapture sesuai ?
//kita gunakan nilai i yang pass by value seperti pada kode ini
func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i) //output tidak urut karena goroutine jalan secara pseudorandom
		}(i) // <- variabel i pass by value
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stop")
}
