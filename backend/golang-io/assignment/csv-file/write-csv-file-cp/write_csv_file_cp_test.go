package main_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	main "github.com/ruang-guru/playground/backend/golang-io/csv-file/write-csv-file-cp"
)

var _ = Describe("WriteToCSV", func() {
	It("should generate CSV file with the correct data", func() {
		err := main.WriteToCSV("./users.csv", [][]string{
			{"name", "age", "country"},
			{"levi", "32", "Indonesia"},
			{"jeff", "30", "USA"},
		})

		Expect(err).ToNot(HaveOccurred())

		file, err := os.Open("./users.csv")
		Expect(err).ToNot(HaveOccurred())

		contents, err := ioutil.ReadAll(file)
		Expect(err).ToNot(HaveOccurred())

		Expect(string(contents)).To(Equal("name,age,country\nlevi,32,Indonesia\njeff,30,USA\n"))

		os.Remove("./users.csv")
	})

})
