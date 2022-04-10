package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Delete", func() {
		When("there is set is not defind element", func() {
			It("should return error", func() {
				set := NewSet()
				_, err := set.Delete("Aditira")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("delete failed!, element Aditira is not in the set"))
			})
		})
		When("there is set is defind element", func() {
			It("should return status true", func() {
				set := NewSet()
				set.Add("Dito")
				status, err := set.Delete("Dito")
				Expect(err).To(BeNil())
				Expect(status).To(BeTrue())
			})
		})
	})
	Describe("Contains", func() {
		When("there is set is not defind element", func() {
			It("should return status false", func() {
				set := NewSet()
				status := set.Contains("Aditira")
				Expect(status).To(BeFalse())
			})
		})
		When("there is set is defind element", func() {
			It("should return status true", func() {
				set := NewSet()
				set.Add("Dito")
				status := set.Contains("Dito")
				Expect(status).To(BeTrue())
			})
		})
	})
})
