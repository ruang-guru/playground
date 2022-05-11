package main

import "fmt"

type PhoneRow struct {
	ID          int // primary key
	CountryCode int
	Number      int
	UserID      int // foreign key
}

type UserRow struct {
	ID   int // primary key
	Name string
	Age  int
}

type PhoneTable []PhoneRow
type UserTable []UserRow

func main() {
	db := make(UserTable, 0)
	db2 := make(PhoneTable, 0)

	db.InsertUser("John", 20)
	db.InsertUser("Gina", 20)

	db2.InsertPhone(62, 1234567890, 1)
	db2.InsertPhone(62, 2345678901, 1)
	db2.InsertPhone(62, 3243434343, 1)

	fmt.Println("Phones:")
	for _, row := range db2 {
		fmt.Println("| UserID: ", db.GetUser(row.UserID).ID, "| Name: ", db.GetUser(row.UserID).Name, "| Age: ", db.GetUser(row.UserID).Age, "| PhoneID: ", row.ID, "| Country Code: ", row.CountryCode, "| Number: ", row.Number, "|")
	}
}

func NewPhoneDB() PhoneTable {
	return make(PhoneTable, 0)
}

func NewUserDB() UserTable {
	return make(UserTable, 0)
}

func (db *PhoneTable) InsertPhone(countryCode int, number int, userID int) {
	(*db) = append(*db, PhoneRow{
		ID:          len(*db) + 1,
		CountryCode: countryCode,
		Number:      number,
		UserID:      userID,
	})
}

func (db *UserTable) InsertUser(name string, age int) {
	(*db) = append(*db, UserRow{
		ID:   len(*db) + 1,
		Name: name,
		Age:  age,
	})
}

func (db *PhoneTable) WherePhone(id int) *PhoneRow {
	// TODO: answer here
}

func (db *UserTable) GetUser(userID int) UserRow {
	var result UserRow
	// TODO: answer here
	return result
}
