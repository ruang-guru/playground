package main

import "sync"

type request struct {
	data     int
	response chan int
}

const maxThroughput = 5
const numOfRequests = 1000

//bagaimana cara membatasi data yang diproses ? hint:buffered channel
func doubleCalculatorWorker(queue chan request, maxThroughput int, maxObservedThroughtputC chan int) {
	// TODO: answer here

	maxObservedThroughtput := 0
	curThroughtput := 0

	//kita akan pelajari mutex di sesi berikutnya
	//secara sederhananya mutex digunakan untuk memastikan hanya ada satu goroutine yang mengakses suatu bagian kode
	mu := &sync.Mutex{}
	for req := range queue {
		// TODO: answer here
		go func(req request) {
			mu.Lock()
			curThroughtput++

			if curThroughtput > maxObservedThroughtput {
				maxObservedThroughtput = curThroughtput
			}
			mu.Unlock()

			data := req.data
			req.response <- data * data

			mu.Lock()
			curThroughtput--
			mu.Unlock()

			// TODO: answer here
		}(req)
	}
	maxObservedThroughtputC <- maxObservedThroughtput
}

func createRequest(queue chan request, result []int) {
	var wg sync.WaitGroup //WaitGroup akan dipelajari pada sesi berikutnya
	//secara simpelnya disini WaitGroup digunakan untuk menunggu goroutine yang dibuat dalam loop ini selesai berjalan
	for i := 0; i < numOfRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			callResult := make(chan int)
			queue <- request{i, callResult}

			result[i] = <-callResult
			close(callResult)
		}(i)
	}
	wg.Wait()
	close(queue) // tutup channel disini, karena fungsi ini yang membuat mengirim ke channel
}
