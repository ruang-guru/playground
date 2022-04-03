package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	fileName := "./original.txt"
	stringToAdd := fmt.Sprint("\nThis is a new string to a line")
	stringOriginal := "I'm old and and im fine"

	createOriginalFile := func() error {
		//membuat file
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
			return err
		}

		//menutup file sebelum fungsi selesai dijalankan
		defer file.Close()

		_, err = file.WriteString(stringOriginal)

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
			return err
		}

		return nil
	}

	BeforeEach(func() {
		err := createOriginalFile()
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
	})

	Describe("AddString", func() {
		It("Should add string to a already existing file", func() {
			data, err := ioutil.ReadFile(fileName)
			Expect(err).To(BeNil())
			Expect(string(data)).To(Equal("I'm old and and im fine"))

			err = AddString(fileName, stringToAdd)
			Expect(err).To(BeNil())

			// read file again
			data, err = ioutil.ReadFile(fileName)
			Expect(err).To(BeNil())
			Expect(string(data)).To(Equal("I'm old and and im fine\nThis is a new string to a line"))
		})
	})
})
