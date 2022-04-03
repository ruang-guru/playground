package main

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	fileName := "write.txt"
	fileData := fmt.Sprint("Write anything here\n" +
		"This program demonstrates reading and writing\n" +
		"operations to a file in Go")

	AfterEach(func() {
		err := os.Remove("./write.txt")
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Writefile", func() {
		It("Should write the string to a file named ./write.txt", func() {

			err := WriteFile(fileName, fileData)
			Expect(err).To(BeNil())

			// Check if file is created
			file, err := os.Stat(fileName)
			Expect(err).To(BeNil())
			Expect(file.Name()).To(Equal(fileName))
			Expect(file.Size()).To(Equal(int64(len(fileData))))
		})
	})
})
