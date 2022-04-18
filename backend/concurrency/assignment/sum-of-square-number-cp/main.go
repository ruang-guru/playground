package main

import "fmt"

//anda diminta untuk mendapatkan hasil pangkat 2 dari angka 1-20
//gunakan fungsi squareWorker untuk menerima dan mengirim data yang sudah dipangkatkan

//fungsi ini digunakan untuk menerima angka lalu mengembalikan hasil pangkat 2 angka tersebut
func squareWorker(workerInput <-chan int, workerOutput chan<- int) {
	for {
		num := <-workerInput      // menerima angka
		workerOutput <- num * num // mengirim hasil
	}
}

func createRequest(workerInput, resultChan chan<- int, workerOutput <-chan int) {
	resultSum := 0 // jumlah dari pemangkatan angka 1-15
	for i := 1; i < 16; i++ {
		var res int
		// TODO: answer here
		resultSum += res
		fmt.Printf("hasil pangkat 2 angka %d adalah: %d\n", i, res) // 0 1 4 9 16
	}
	fmt.Printf("total penjumlahan : %d\n", resultSum)
	resultChan <- resultSum
}
