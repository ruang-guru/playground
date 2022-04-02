package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

//jalankan dari dalam folder write

//membuka file jika sudah ada atau membuat baru jika belum ada
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
	newData := [][]string{{"nama", "nilai", "kelas"},
		{"andi", "90", "A"},
		{"anda", "92", "A"}}

	err := saveCSV("siswaA", newData)
	if err != nil {
		fmt.Println("error : ", err)
	}

	// reset() //uncomment this to remove siswaA.csv
}

//data akan masuk ke path berikut : ../data/siswaA.csv

func reset() {
	os.Remove("../data/siswaA.csv")
}
