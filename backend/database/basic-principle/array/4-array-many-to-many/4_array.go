package main

import "fmt"

type PhoneRow struct {
	ID          int // primary key
	CountryCode int
	Number      int
}
type UserRow struct {
	ID   int // primary key
	Name string
	Age  int
}
type UserPhoneRow struct {
	UserID  int // primary key
	PhoneID int // primary key
}

type PhoneTable []PhoneRow
type UserTable []UserRow
type UserPhoneTable []UserPhoneRow

func main() {
	db := make(UserTable, 0)
	db2 := make(PhoneTable, 0)
	db3 := make(UserPhoneTable, 0)

	db.InsertUser("John", 20)
	db.InsertUser("Gina", 20)

	db2.InsertPhone(62, 1234567890)
	db2.InsertPhone(62, 2345678901)

	db3.InsertUserPhone(1, 1)
	db3.InsertUserPhone(1, 2)
	db3.InsertUserPhone(2, 1)
	db3.InsertUserPhone(2, 2)

	fmt.Println("Phones:")
	for _, row := range db2 {
		fmt.Println("| ID: ", row.ID, "| Country Code: ", row.CountryCode, "| Number: ", row.Number, "|")
	}
	fmt.Println("Users:")
	for _, row := range db {
		fmt.Println("| ID: ", row.ID, "| Name: ", row.Name, "| Age: ", row.Age, "|")
	}
	fmt.Println("UserPhone:")
	for _, row := range db3 {
		fmt.Println("| User ID: ", row.UserID, "| Phone ID: ", row.PhoneID, "|")
	}
}

func (db *PhoneTable) InsertPhone(countryCode int, number int) {
	(*db) = append(*db, PhoneRow{
		ID:          len(*db) + 1,
		CountryCode: countryCode,
		Number:      number,
	})
}

func (db *UserTable) InsertUser(name string, age int) {
	(*db) = append(*db, UserRow{
		ID:   len(*db) + 1,
		Name: name,
		Age:  age,
	})
}

func (db *UserPhoneTable) InsertUserPhone(userID int, phoneID int) {
	(*db) = append(*db, UserPhoneRow{
		UserID:  userID,
		PhoneID: phoneID,
	})
}
