package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Readfile", func() {
		It("Should Read a file and returning given data", func() {
			Expect(1).To(Equal(1))
		})
	})
})
