package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Count Words", func() {
		When("input a sentence with a few words", func() {
			It("should return return word count value", func() {
				data := CountWords("Andi suka bermain bola")
				Expect(data).To(Equal(4))
			})
		})
	})

	Describe("MostWordsFound", func() {
		When("when input list with each line having multiple sentences", func() {
			It("should return value of the most sentences from the list entered", func() {
				data := MostWordsFound([]string{"Andi suka bermain bola", "Saya sedang belajar struktur data", "Terima kasih"})
				Expect(data).To(Equal(5))
			})
		})
	})
})
