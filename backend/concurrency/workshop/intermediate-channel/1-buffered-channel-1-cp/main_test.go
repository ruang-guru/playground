package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Buffered Channel", func() {
	When("sending", func() {
		It("won't block, until the buffer full", func() {
			output := make(chan int)
			go testBuffer(output)
			Expect(<-output).To(Equal(6))
		})
	})
})
