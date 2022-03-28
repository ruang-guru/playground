package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Custom Error", func() {
	peopleAge := make(map[string]int)
	BeforeEach(func() {
		peopleAge = map[string]int{
			"John": 20,
			"Raz":  8,
			"Tony": -1,
		}
	})

	Describe("Get Age with invalid data", func() {
		It("returns ErrorInvalidData", func() {
			_, err := GetAge(peopleAge, "Tony")
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("500"))
			Expect(err.Error()).Should(ContainSubstring("error invalid data"))
		})
	})
	Describe("Get Age with no data", func() {
		It("returns error data not found", func() {
			_, err := GetAge(peopleAge, "Roger")
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("Data not found"))
		})
	})
})
