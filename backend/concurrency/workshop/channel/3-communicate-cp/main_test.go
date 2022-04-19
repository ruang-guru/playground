package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//gunakan channel untuk mengirimkan data
//note: fungsi main atau test juga berjalan dengan goroutine
var _ = Describe("Channel", func() {
	It("can communicate between goroutine", func() {
		output := make(chan string)
		go communicate(output)
		Expect(<-output).To(Equal("hello steve"))
	})
})
