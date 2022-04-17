package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Fibonacci", func() {
		When("input number more than 1", func() {
			It("should return the fibonacci number", func() {
				Expect(FibonacciRecursion(5)).To(Equal(5))
				Expect(FibonacciRecursion(6)).To(Equal(8))
				Expect(FibonacciRecursion(7)).To(Equal(13))
				Expect(FibonacciRecursion(8)).To(Equal(21))
				Expect(FibonacciRecursion(9)).To(Equal(34))
				Expect(FibonacciRecursion(10)).To(Equal(55))
			})
		})
	})
})
