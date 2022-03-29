package main

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Is", func() {
	peopleAge := make(map[string]int)
	BeforeEach(func() {
		peopleAge = map[string]int{
			"Roger":  60,
			"Kamala": 5,
			"Peter":  20,
		}
	})

	Describe("Get IsEligibleToVaccine with error in GetAge", func() {
		It("returns wrapped error", func() {
			_, err := IsEligibleToVaccine(peopleAge, "Tony")
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("error in IsEligibleToVaccine"))
			Expect(errors.Is(err, ErrDataNotFound)).To(BeTrue())
		})
	})
})
