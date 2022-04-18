package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//Gunakan channel untuk melakukan blocking
//agar goroutine dapat berjalan
var _ = Describe("Channel", func() {
	When("receiving", func() {
		It("can block", func() {
			output := make(chan int)
			go receiveBlock(output)
			Expect(<-output).To(Equal(1))
		})
	})
})
