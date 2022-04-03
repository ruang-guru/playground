package main_test

import (
	"io/ioutil"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	main "github.com/ruang-guru/playground/backend/golang-io/json-file/write-json-file-cp"
)

var _ = Describe("WriteToJson", func() {
	It("should generate Json file with the correct data", func() {
		err := main.WriteToJson("./users.json", []main.User{
			{
				Name:    "levi",
				Age:     32,
				Country: "Indonesia",
			},
			{
				Name:    "jeff",
				Age:     30,
				Country: "USA",
			},
		})
		Expect(err).ToNot(HaveOccurred())

		file, err := os.Open("./users.json")
		Expect(err).ToNot(HaveOccurred())

		contents, err := ioutil.ReadAll(file)
		Expect(err).ToNot(HaveOccurred())

		Expect(strings.TrimSpace(string(contents))).To(Equal("[{\"Name\":\"levi\",\"Age\":32,\"Country\":\"Indonesia\"},{\"Name\":\"jeff\",\"Age\":30,\"Country\":\"USA\"}]"))

		os.Remove("./users.json")
	})
})
