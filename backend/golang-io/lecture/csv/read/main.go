package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

//jalankan dari dalam folder read

//membuka file
func openFile(csvName string) (*os.File, error) {
	path, err := filepath.Abs("../data/" + csvName + ".csv")

	if err != nil {
		return nil, err
	}
	studentsFile, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
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

func main() {
	data, err := loadCSV("siswaB")
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println(data)
}
