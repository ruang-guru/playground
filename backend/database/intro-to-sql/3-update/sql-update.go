package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db}
}

func (r *StudentRepository) UpdateStudent(student Student) error {
	// stmt dibawah ini adalah syntax untuk mengupdate data
	// dalam tabel student dengan nim yang sama dengan nim yang diberikan
	// data yang diupdate adalah nama dan umur
	stmt := `
		UPDATE students
		SET name = ?, age = ?
		WHERE nim = ?
	`
	_, err := r.db.Exec(stmt, student.Name, student.Age, student.NIM)
	if err != nil {
		return err
	}

	return nil
}

func (r *StudentRepository) FetchStudents() ([]model.Student, error) {
	// stmt dibawah ini adalah syntax untuk mengambil data
	// dari tabel student
	sqlStmt := `SELECT id, name, age, nim FROM students;`

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	students := []model.Student{}
	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.NIM)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
