package main

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteToCSV(filePath string, rows [][]string) error {
	// TODO: answer here
	return nil
}

func main() {

	err := WriteToCSV("./users.csv", [][]string{
		{"name", "age", "country"},
		{"levi", "32", "Indonesia"},
		{"jeff", "30", "USA"},
	})

	if err != nil {
		log.Fatal(err)
	}
}
