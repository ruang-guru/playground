package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("mutex", func() {
	It("can help solve race condition problem", func() {
		output := make(chan int)
		go counter(output)
		Expect(<-output).To(Equal(1000))
	})

})
