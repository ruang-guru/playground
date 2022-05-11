package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{db}
}

func (r *StudentRepo) CreateStudent(student model.Student) (int64, error) {
	// sql statement yang digunakan untuk menambahkan data student baru
	sqlStmt := `INSERT INTO students (name, age, nim)
		VALUES (?, ?, ?);`

	// membuat prepared statement dengan sql statement yang sudah dibuat
	// dan mengirimkan parameter yang dibutuhkan
	// untuk menambahkan data student baru
	result, err := r.db.Exec(sqlStmt, student.Name, student.Age, student.NIM)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *StudentRepo) FetchStudents() ([]model.Student, error) {
	// sql statement yang digunakan untuk mengambil data students dari database
	sqlStmt := `SELECT id, name, age, nim FROM students;`

	// membuat prepared statement dengan sql statement yang sudah dibuat
	// melakukan eksekusi query dan mengambil data students
	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	// membuat slice untuk menampung data students
	students := []model.Student{}

	// melooping data students yang diterima dari database
	// dan menambahkan data students ke dalam slice
	// rows.Next() berfungsi untuk mengambil data student berdasarkan urutan
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
