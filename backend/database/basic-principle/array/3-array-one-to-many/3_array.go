package main

import "fmt"

type PhoneRow struct {
	ID          int // primary key
	CountryCode int
	Number      int
}

type UserRow struct {
	ID       int // primary key
	Name     string
	Age      int
	PhoneIDs []int // foreign key
}

type UserTable []UserRow
type PhoneTable []PhoneRow

func main() {
	db := make(UserTable, 0)
	db2 := make(PhoneTable, 0)

	db.InsertUser("John", 20, []int{1, 2})
	db.InsertUser("Gina", 20, []int{1, 2, 3})

	db2.InsertPhone(62, 1234567890)
	db2.InsertPhone(62, 2345678901)
	db2.InsertPhone(62, 3243434343)

	fmt.Println("Users:")
	for _, row := range db {
		for _, phoneID := range row.PhoneIDs {
			fmt.Println("| ID: ", row.ID, "| Name: ", row.Name, "| Age: ", row.Age, "| Phone Data: ", db2.GetPhone(phoneID), "|")
		}
	}
}

func (db *UserTable) InsertUser(name string, age int, phoneIDs []int) {
	(*db) = append(*db, UserRow{
		ID:       len(*db) + 1,
		Name:     name,
		Age:      age,
		PhoneIDs: phoneIDs,
	})
}

func (db *PhoneTable) InsertPhone(countryCode int, number int) {
	(*db) = append(*db, PhoneRow{
		ID:          len(*db) + 1,
		CountryCode: countryCode,
		Number:      number,
	})
}

func (db *PhoneTable) GetPhone(phoneID int) PhoneRow {
	var result PhoneRow
	for _, row := range *db {
		if row.ID == phoneID {
			result = row
		}
	}
	return result
}
