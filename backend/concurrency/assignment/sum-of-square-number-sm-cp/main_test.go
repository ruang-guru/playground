package main

import (
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goroutine", func() {
	When("using shared memory model", func() {
		It("can be used with waitgroup and mutex", func() {
			workerInput := make(chan int)
			workerOutput := make(chan int)
			var wg sync.WaitGroup
			wg.Add(1)
			go squareWorker(workerInput, workerOutput)
			go createRequest(workerInput, workerOutput, &wg)
			wg.Wait()
			mu.Lock()
			Expect(result).To(Equal(328350))
			mu.Unlock()
		})
	})
})
