package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("CSVToMap", func() {
		It("Should return a map with structure like this `question:answer`", func() {
			data := make(map[string]string)
			data, err := CSVToMap(data, "questions.csv")
			Expect(err).To(BeNil())

			Expect(data["Nama ibukota indonesia?"]).To(Equal("Jakarta"))
			Expect(data["Siapakah presiden pertama Indonesia?"]).To(Equal("Soekarno"))
		})
	})
})
