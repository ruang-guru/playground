package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sentinel Error", func() {
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
			Expect(err.Error()).Should(ContainSubstring("error age is invalid, less than 0"))
			Expect(err).To(MatchError(ErrInvalidAge))
		})
	})
	Describe("Get Age with no data", func() {
		It("returns error data not found", func() {
			_, err := GetAge(peopleAge, "Roger")
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("error data not found"))
			Expect(err).To(MatchError(ErrDataNotFound))
		})
	})
})
