package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	stdscore "github.com/ruang-guru/playground/backend/basic-golang/template/1-get-struct-value-cp-advance"
)

var _ = Describe("Student", func() {

	Describe("Calculate", func() {
		It("calculate score to get score average", func() {
			std := stdscore.NewStudent("Rogu", []float64{10, 11, 12})
			result := std.CalculateScore(std.Scores)
			Expect(result).To(Equal(float64(11)))
		})
	})

	Describe("Generate", func() {
		It("generate student score template", func() {
			std := stdscore.NewStudent("Rogu", []float64{10, 11, 12})
			result := std.GenerateStudentTemplate()
			Expect(result).To(Equal("Hello Rogu, Nilai rata-rata kamu 11"))
		})
	})
})
