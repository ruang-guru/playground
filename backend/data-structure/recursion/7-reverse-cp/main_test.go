package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Reverse", func() {
		When("input list of letters by forming the word upside down", func() {
			It("should return the reversed word", func() {
				str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
				s := Reverse(str, len(str)-1)
				Expect(s).To(Equal("INDONESIA"))
			})
		})
	})
})
