package main

import "fmt"

type StudentRow struct {
	ID       int // primary key
	Name     string
	Age      int
	SchoolID int // foreign key
}

type SchoolRow struct {
	ID      int // primary key
	Name    string
	Address string
}

type StudentTable []StudentRow
type SchoolTable []SchoolRow

func main() {
	db := make(StudentTable, 0)
	db2 := make(SchoolTable, 0)

	db.InsertStudent("John", 20, 1)
	db.InsertStudent("Jane", 21, 1)
	db.InsertStudent("Gina", 20, 2)

	db2.InsertSchool("SDIT", "Jakarta")
	db2.InsertSchool("SDIT", "Bandung")

	println("Student:")
	for _, row := range db {
		fmt.Println("| ID:", row.ID, " | Student:", row.Name, " | Age:", row.Age, " | SchoolID:", db2.GetSchool(row.SchoolID), " |")
	}
}

func NewSchoolDB() SchoolTable {
	return make(SchoolTable, 0)
}

func NewStudentDB() StudentTable {
	return make(StudentTable, 0)
}

func (db *StudentTable) InsertStudent(name string, age int, schoolID int) {
	(*db) = append(*db, StudentRow{
		ID:       len(*db) + 1,
		Name:     name,
		Age:      age,
		SchoolID: schoolID,
	})
}

func (db *SchoolTable) InsertSchool(name string, address string) {
	(*db) = append(*db, SchoolRow{
		ID:      len(*db) + 1,
		Name:    name,
		Address: address,
	})
}

func (db *StudentTable) WhereStudent(id int) *StudentRow {
	// TODO: answer here
}

func (db *SchoolTable) GetSchool(schoolID int) SchoolRow {
	// TODO: answer here
}
