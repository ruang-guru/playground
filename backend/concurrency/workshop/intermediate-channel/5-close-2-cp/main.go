package main

import (
	"fmt"
	"time"
)

//numberWorker digunakan untuk mengirim data ke channel output
func numberWorker(output chan int) {
	for i := 0; i < 100; i++ {
		//kirim ke channel
		// TODO: answer here

	}
	close(output)
	//pastikan tulisan ini berhasil ditampilkan
	fmt.Println("finished sending data. channel closed")
}

func receiver(result chan int) {
	//buat channel dengan ukuran yang sesuai dengan data-
	//yang akan dikirim numWorker
	output:=make(chan int) // TODO: replace this
	sum := 0

	go numberWorker(output)
	time.Sleep(100 * time.Millisecond)
	for number := range output {
		square := number * number
		sum += square
	}
	result <- sum
}
