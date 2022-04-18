package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Count Digits", func() {
		When("input some number", func() {
			It("should return the number of digits", func() {
				Expect(CountDigits(1996)).To(Equal(4))
			})
		})
	})
})
