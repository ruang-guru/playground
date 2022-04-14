package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Selection Search", func() {
		When("input list is not sorted", func() {
			It("should return a sorted list", func() {
				sample := []int{3, 4, 5, 2, 1}
				arrSorted := SelectionSort(sample)
				Expect(arrSorted).To(Equal([]int{1, 2, 3, 4, 5}))
			})
		})
	})
})
