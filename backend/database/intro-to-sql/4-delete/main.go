package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	studentRepo := NewStudentRepository(db)

	err = studentRepo.DeleteStudentByNIM("12345")
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete Success")

	// Use FetchStudents to get all students
	students, err := studentRepo.FetchStudents()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID\tName\tNIM\tAge\n")
	for _, student := range students {
		fmt.Printf("%d\t%s\t%s\t%d\n", student.ID, student.Name, student.NIM, student.Age)
	}
}
