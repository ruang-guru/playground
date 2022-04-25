package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sum All", func() {
	It("Sum Asc From 0", func() {
		res := SumAllCorrect([]int{0, 1, 2, 3, 4, 5})
		Expect(res).To(Equal(15))
	})

	It("Sum Random Number", func() {
		res := SumAllCorrect([]int{12, 5, 10, 1, 3})
		Expect(res).To(Equal(31))
	})
})
