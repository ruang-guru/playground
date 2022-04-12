package main

import (
	"encoding/csv"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("WriteTo(csvFileName, data)", func() {
		It("should write to csv file, based on data given", func() {
			csvFileName := "menu.csv"
			data := []Menu{
				{"Pizza", 100},
				{"Burger", 200},
				{"Coffee", 300},
				{"Tea", 400},
				{"Sandwich", 500},
			}

			err := WriteToCSV(csvFileName, data)
			Expect(err).To(BeNil())

			csvFile, err := os.Open(csvFileName)
			Expect(err).To(BeNil())
			defer csvFile.Close()

			csvReader := csv.NewReader(csvFile)
			records, err := csvReader.ReadAll()
			Expect(err).To(BeNil())
			Expect(len(records)).To(Equal(5))
			Expect(records[0]).To(Equal([]string{"Pizza", "100"}))
			Expect(records[1]).To(Equal([]string{"Burger", "200"}))
			Expect(records[2]).To(Equal([]string{"Coffee", "300"}))
			Expect(records[3]).To(Equal([]string{"Tea", "400"}))
			Expect(records[4]).To(Equal([]string{"Sandwich", "500"}))
		})
	})
})
