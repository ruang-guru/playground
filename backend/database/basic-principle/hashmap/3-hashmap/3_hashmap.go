package main

import (
	"fmt"
)

type PrimaryKey int

type SecondaryKey string

type SalesRow struct {
	ID       PrimaryKey   //ID must be unique
	Customer SecondaryKey //Name can be non-unique
	Date     string
	Note     string
}

type IndexByID map[PrimaryKey]SalesRow

type IndexByName map[SecondaryKey][]PrimaryKey

type SalesDB struct {
	ByID   IndexByID
	ByName IndexByName
}

func main() {
	db := SalesDB{
		ByID:   make(IndexByID),
		ByName: make(IndexByName),
	}
	db.Insert("John", "2019-01-01", "Note 1")
	db.Insert("John", "2019-01-02", "Note 2")
	db.Insert("Gina", "2019-01-03", "Note 3")
	db.Insert("Gina", "2019-01-04", "Note 4")
	db.Insert("Gina", "2019-01-05", "Note 5")
	db.Insert("Junar", "2019-01-06", "Note 6")
	db.Insert("Bob", "2019-01-07", "Note 7")
	db.Insert("Ayumi", "2019-01-08", "Note 8")

	// All data:
	fmt.Println("ByID:")
	for _, row := range db.ByID {
		fmt.Println(" -> ", row)
	}
	fmt.Println("ByName:")
	for _, row := range db.ByName {
		fmt.Println(" -> ", row)
	}

	fmt.Println("|========================================================|")

	fmt.Println("1. Where by ID: 8", db.WhereByID(8))

	fmt.Println("2. Where ID is large: 5")
	for _, row := range db.WhereIDIsLargeThan(5) {
		fmt.Println(" -> ", row)
	}
}

func (db *SalesDB) Insert(customer string, date string, note string) {
	db.ByID[PrimaryKey(len(db.ByID))+1] = SalesRow{
		ID:       PrimaryKey(len(db.ByID)) + 1,
		Customer: SecondaryKey(customer),
		Date:     date,
		Note:     note,
	}
	db.ByName[SecondaryKey(customer)] = append(db.ByName[SecondaryKey(customer)], PrimaryKey(len(db.ByID)))
}

func (db *SalesDB) WhereByID(id PrimaryKey) *SalesRow {
	m, ok := db.ByID[id]
	if !ok {
		return nil
	}
	return &m
}

func (db *SalesDB) WhereIDIsLargeThan(id PrimaryKey) []SalesRow {
	rows := make([]SalesRow, 0)
	for _, row := range db.ByID {
		if row.ID > id {
			rows = append(rows, row)
		}
	}
	return rows
}
