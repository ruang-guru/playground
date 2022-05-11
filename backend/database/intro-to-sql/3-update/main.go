package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Student struct {
	ID   int    `db:"id"`
	NIM  string `db:"nim"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	studentRepo := NewStudentRepository(db)

	student := Student{
		NIM:  "12345",
		Name: "Johns",
		Age:  25,
	}

	err = studentRepo.UpdateStudent(student)
	if err != nil {
		panic(err)
	}

	fmt.Println("Student updated")

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
