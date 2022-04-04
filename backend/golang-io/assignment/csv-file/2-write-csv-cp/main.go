package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Menu struct {
	Name  string
	Price int
}

func WriteToCSV(fileName string, records []Menu) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// TODO: answer here
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, val := range records {
		row := []string{val.Name, strconv.Itoa(val.Price)}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	return nil
}
