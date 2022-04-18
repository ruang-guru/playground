package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Channel", func() {
	When("used as function parameter", func() {
		It("won't have behavior change", func() {
			output := make(chan string)

			a := "andy"
			b := "mandy"
			go start(a, b, output)

			Expect(<-output).To(Equal("halo andy dan mandy"))
		})
	})
})
