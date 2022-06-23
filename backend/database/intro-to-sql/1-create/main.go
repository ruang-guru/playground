package main

import (
	"database/sql"
	"fmt"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	studentRepo := NewStudentRepo(db)

	resultID, err := studentRepo.CreateStudent(model.Student{
		Name: "Robert",
		NIM:  "L20019",
		Age:  20,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("ID:", resultID)

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
