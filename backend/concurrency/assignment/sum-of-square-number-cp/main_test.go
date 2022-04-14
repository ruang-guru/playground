package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Channel", func() {
	It("can communicate between goroutine", func() {
		workerInput := make(chan int)
		workerOutput := make(chan int)
		resultChan := make(chan int)
		go squareWorker(workerInput, workerOutput)
		go createRequest(workerInput, resultChan, workerOutput)
		Expect(<-resultChan).To(Equal(1240))
	})
})
