package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	main "github.com/ruang-guru/playground/backend/unittest"
)

var _ = Describe("Main", func() {
	Describe("Add", func() {
		When("both numbers are positive", func() {
			It("returns a positive number", func() {
				Expect(main.Sum(2, 2)).To(BeNumerically(">", 0))
			})
		})

		When("both numbers are negative", func() {
			It("returns a negative number", func() {
				Expect(main.Sum(-2, -2)).To(BeNumerically("<", 0))
			})
		})

		When("both numbers are zero", func() {
			It("returns zero", func() {
				Expect(main.Sum(0, 0)).To(BeZero())
			})
		})

		It("add two numbers", func() {
			Expect(main.Sum(2, 2)).To(Equal(4))
		})
	})
})
