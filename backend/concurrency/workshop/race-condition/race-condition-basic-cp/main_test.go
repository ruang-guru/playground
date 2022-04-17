package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("race condition", func() {
	When("happening", func() {
		It("can be fixed by using channel", func() {
			output := make(chan int)
			go counter(output)
			Expect(<-output).To(Equal(1000))
		})
	})
})
