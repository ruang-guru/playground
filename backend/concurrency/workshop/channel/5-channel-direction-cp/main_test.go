package main

import (
	"runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Channel", func() {
	When("used as function parameter", func() {
		It("direction can be specified", func() {
			runtime.GOMAXPROCS(1)
			input := make(chan int)
			output := make(chan int)
			go numberGenerator(10, input)
			go squareWorker(input, output)
			for i := 1; i <= 10; i++ {
				Eventually(output, "100ms").Should(Receive(Equal(i * i)))
			}
		})
	})
})
