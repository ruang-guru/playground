package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("channel", func() {
	When("using for ... range channel", func() {
		It("can loop while channel still active", func() {
			output := make(chan int)
			go receiver(output)
			for i := 0; i < 10; i++ {
				Eventually(output, "100ms").Should(Receive(Equal(i * i)))
			}
		})
	})
})
