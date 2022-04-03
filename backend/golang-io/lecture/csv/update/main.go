package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

//jalankan dari dalam folder update

//membuka file
func openFile(csvName string) (*os.File, error) {
	path, err := filepath.Abs("../data/" + csvName + ".csv")
	if err != nil {
		return nil, err
	}
	studentsFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return studentsFile, nil
}

func loadCSV(dbName string) ([][]string, error) {
	f, err := openFile(dbName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func saveCSV(dbName string, records [][]string) error {
	f, err := openFile(dbName)
	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(records)
	if err != nil {
		return err
	}

	w.Flush()
	return nil
}

func main() {
	dbName := "siswaB"
	data, err := loadCSV(dbName)
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Printf("sebelum update : %s\n", data)
	newStudent := []string{"Raam", "90", "B"}
	data = append(data, newStudent)
	err = saveCSV(dbName, data)
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Printf("sesudah update : %s\n", data)
	reset(dbName) // comment this to show change in csv file
}

func reset(dbName string) {
	initialData := [][]string{{"nama", "nilai", "kelas"},
		{"aba", "90", "B"},
		{"abi", "92", "B"},
		{"abo", "80", "B"},
		{"abu", "88", "B"},
		{"abe", "85", "B"}}
	saveCSV(dbName, initialData) //initial data
}
