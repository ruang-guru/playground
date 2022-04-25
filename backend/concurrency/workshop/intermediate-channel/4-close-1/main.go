package main

import (
	"fmt"
	"time"
)

func main() {
	output := make(chan string, 10)

	go sender(output)
	for msg := range output {
		fmt.Println(msg)
	}
}

func sender(output chan string) {
	for i := 0; i < 10; i++ {
		output <- "every 500 millisecond"
		time.Sleep(500 * time.Millisecond)
	}
	close(output)
}

//jika menggunakan for ... range channel
//kita perlu menutup channel tersebut saat sudah tidak ada data yang dikirim
//yang menutup channel sebaiknya adalah pengirim data (disini fungsi sender)
