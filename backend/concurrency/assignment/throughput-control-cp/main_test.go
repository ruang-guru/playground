package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//dengan buffered channel kita bisa membatasi throughput yang diterima
//hal ini bisa dicapai karena pada buffered channel blocking terjadi ketika buffer channel penuh
var _ = Describe("buffered channel", func() {
	It("can be used to control the maximum throughtput", func() {
		maxObservedThroughtputC := make(chan int)
		queue := make(chan request)
		go doubleCalculatorWorker(queue, maxThroughput, maxObservedThroughtputC)

		result := make([]int, numOfRequests)
		createRequest(queue, result)
		for i := 0; i < numOfRequests; i++ {
			Expect(result[i]).To(Equal(i * i))
		}

		Expect(<-maxObservedThroughtputC).To(BeNumerically("<=", maxThroughput))
	})
})
