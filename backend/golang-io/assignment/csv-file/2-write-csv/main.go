package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	ID  string
	Age int
}

func main() {
	csvFile, err := os.Create("employee.csv")
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", "employee.csv", err)
	}

	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Write some records to file
	records := []Employee{
		{"E01", 25},
		{"E02", 26},
		{"E03", 24},
		{"E04", 26},
	}

	// Using Write
	for _, record := range records {
		row := []string{record.ID, strconv.Itoa(record.Age)}
		if err := csvWriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	// Uncomment this to try it out, Comment Above
	// // Using WriteAll
	// var data [][]string
	// for _, record := range records {
	// 	row := []string{record.ID, strconv.Itoa(record.Age)}
	// 	data = append(data, row)
	// }
	// csvWriter.WriteAll(data)
}
