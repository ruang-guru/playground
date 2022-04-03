package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Readfile", func() {
		It("Should Read a file and returning given data", func() {
			fileData, err := ReadFile("./read.txt")
			Expect(err).To(BeNil())
			isiFile := string(fileData.Data)
			Expect(fileData.Name).To(Equal("./read.txt"))
			Expect(isiFile).To(Equal("Halo teman teman MBKM"))
			Expect(fileData.Size).To(Equal(len(fileData.Data)))
		})
	})
})
