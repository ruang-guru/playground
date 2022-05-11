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

func (r *StudentRepository) DeleteStudentByNIM(nim string) error {
	// statement dibawah ini akan menghapus data student dengan nim yang diberikan
	statement := `DELETE FROM students WHERE nim = ?`
	_, err := r.db.Exec(statement, nim)
	return err
}

func (r *StudentRepository) FetchStudents() ([]model.Student, error) {
	// statement dibawah ini akan mengambil data student
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
