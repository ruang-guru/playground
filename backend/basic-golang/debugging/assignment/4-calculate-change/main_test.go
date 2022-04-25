package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate Change", func() {
	Describe("Below Discount Rate", func() {
		It("Normal Payment Price", func() {
			res := CalculateChangeCorrect(5000, 1000)
			Expect(res).To(Equal(float64(4000)))
		})

		It("Normal Payment Price No Change", func() {
			res := CalculateChangeCorrect(5000, 5000)
			Expect(res).To(Equal(float64(0)))
		})

		It("Payment Below Price", func() {
			res := CalculateChangeCorrect(1000, 5000)
			Expect(res).To(Equal(float64(-1)))
		})
	})

	Describe("With Discount Rate", func() {
		It("Normal Payment Price", func() {
			res := CalculateChangeCorrect(100000, 100000)
			Expect(res).To(Equal(float64(5000)))
		})

		It("Payment Below Price", func() {
			res := CalculateChangeCorrect(90000, 100000)
			Expect(res).To(Equal(float64(-1)))
		})

		It("Payment Higher than Price After Discount", func() {
			res := CalculateChangeCorrect(98000, 100000)
			Expect(res).To(Equal(float64(3000)))
		})
	})
})
