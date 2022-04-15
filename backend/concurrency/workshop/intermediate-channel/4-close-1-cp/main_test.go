package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("channel", func() {
	When("closed", func() {
		It("can stop for ... range channel", func() {
			output := make(chan int)
			go receiver(output)
			Eventually(output).Should(BeClosed())
		})
	})
})
