package db

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

type CsvDB struct{}

func NewCsvDB() *CsvDB {
	return &CsvDB{}
}

func openFile(dbName DBName) (*os.File, error) {
	path, err := filepath.Abs("./data/" + dbName + ".csv")
	usersFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return usersFile, nil
}

func (db *CsvDB) Load(dbName DBName) (Rows, error) {
	f, err := openFile(dbName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()
	return rows, nil
}

func (db *CsvDB) Save(dbName DBName, rows Rows) error {
	f, err := openFile(dbName)
	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	w.Flush()
	return nil
}

func (db *CsvDB) Delete(dbName DBName) error {
	err := os.Truncate("./data/"+dbName+".csv", 0)
	if err != nil {
		return err

	}
	return nil
}
