package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var result int

//fungsi ini digunakan untuk menerima angka lalu mengembalikan hasil pangkat 2 angka tersebut
func squareWorker(workerInput <-chan int, workerOutput chan<- int) {
	for {
		num := <-workerInput      // menerima angka
		workerOutput <- num * num // mengirim hasil
	}
}

func createRequest(workerInput chan<- int, workerOutput <-chan int, wg *sync.WaitGroup) {
	for i := 1; i < 100; i++ {
		// TODO: answer here
		wg.Add(1)
		go func(i int) {
			// TODO: answer here
			workerInput <- i
			var res int

			// TODO: answer here
			res = <-workerOutput
			//tambahkan res ke result. Selain itu gunakan juga sesuatu yang menghindari data race
			// TODO: answer here
			result += res
			wg.Done()
			fmt.Println(res)
		}(i)
	}
	wg.Done()
}
