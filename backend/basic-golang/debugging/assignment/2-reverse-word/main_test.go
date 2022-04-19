package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reverse Word", func() {
	It("Palindrome Word", func() {
		res := ReverseWordCorrect("katak")
		Expect(res).To(Equal("katak"))
	})

	It("Normal Word", func() {
		res := ReverseWordCorrect("ruangguru")
		Expect(res).To(Equal("uruggnaur"))
	})
})
