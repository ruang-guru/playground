package main

import (
	"encoding/csv"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Domicile  string
}

func CSVReadAll() ([][]string, error) {
	f, err := os.Open("people.csv")
	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip baris pertama (nama kolom)
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	data, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return data, nil
}
