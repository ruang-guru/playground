package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("channel", func() {
	When("closed", func() {
		It("can still give the data inside", func() {
			result := make(chan int)
			go receiver(result)
			Expect(<-result).To(Equal(328350))
		})
	})
})

//walaupun channel sudah di close, data dalam channel masih dapat diterima
