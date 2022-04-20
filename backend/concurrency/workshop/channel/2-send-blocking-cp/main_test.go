package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//Gunakan channel untuk melakukan blocking
//agar goroutine yang dibuat dapat berjalan
var _ = Describe("Channel", func() {
	When("sending", func() {
		It("can block", func() {
			output := make(chan bool)
			go sendBlock(output)
			Expect(<-output).To(Equal(true))
		})
	})
})
