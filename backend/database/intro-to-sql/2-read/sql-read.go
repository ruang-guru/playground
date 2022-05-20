package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepo {
	return &StudentRepo{db}
}

func (r *StudentRepo) FetchStudentByNIM(nim string) (*model.Student, error) {
	// sql statement dibawah ini akan mengambil data dari tabel student dengan nim yang diberikan
	sqlStmt := `
	SELECT 
	    id, name, age, nim
    FROM 
        students 
    WHERE nim = ?`

	// queryrow untuk mengambil data dari tabel student dengan nim yang diberikan
	row := r.db.QueryRow(sqlStmt, nim)
	// row.Scan digunakan untuk mengambil data dari tabel student dengan nim yang diberikan
	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Age, &student.NIM)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepo) FetchStudents() ([]model.Student, error) {
	// sql statement dibawah ini akan mengambil data dari tabel students
	sqlStmt := `
	SELECT 
	    id, name, age, nim
	FROM 
		students`

	// Query untuk mengambil data dari tabel students
	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	// slice students untuk menampung data dari tabel students
	var students []model.Student
	// rows.Scan digunakan untuk mengambil data dari tabel students dan dimasukkan ke dalam slice students
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
