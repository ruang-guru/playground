package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Least Common Multiple (LCM)", func() {
		When("input some number to calculate LCM", func() {
			It("should return the least common multiple", func() {
				Expect(LCM(2, 3)).To(Equal(6))
				Expect(LCM(10, 15, 20)).To(Equal(60))
				Expect(LCM(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)).To(Equal(2520))
			})
		})
	})
})
