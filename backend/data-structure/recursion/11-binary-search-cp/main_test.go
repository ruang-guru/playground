package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Binnary Search", func() {
		When("binary search is found", func() {
			It("should return 1", func() {
				numList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				key := int64(9)
				result := BinarySearch(numList, key)
				Expect(result).To(Equal(1))
			})
		})
		When("binary search is not found", func() {
			It("should return 0", func() {
				numList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				key := int64(12)
				result := BinarySearch(numList, key)
				Expect(result).To(Equal(0))
			})
		})
	})
})
