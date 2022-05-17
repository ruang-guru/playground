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

	studenRepo := NewStudentRepository(db)

	student, err := studenRepo.FetchStudentByNIM("12345")
	if err != nil {
		panic(err)
	}

	// Print the student
	fmt.Printf("ID\tName\tNIM\tAge\n")
	fmt.Printf("%d\t%s\t%s\t%d\n", student.ID, student.Name, student.NIM, student.Age)

	students, err := studenRepo.FetchStudents()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")

	// Print the students
	fmt.Printf("ID\tName\tNIM\tAge\n")
	for _, student := range students {
		fmt.Printf("%d\t%s\t%s\t%d\n", student.ID, student.Name, student.NIM, student.Age)
	}
}
