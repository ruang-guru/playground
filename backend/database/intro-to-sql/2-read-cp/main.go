package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

func (r *EmployeeRepository) FetchEmployeeByNIK(nik string) (*model.Employee, error) {
	var sqlStmt string

	// Task:
	// buat query untuk mengambil data employee berdasarkan nik
	// simpan query ke variable sqlStmt
	// lihat model.Employee untuk field yang diambil dari database
	// data diambil dengan nik tertentu dari parameter nik

	// TODO: answer here

	row := r.db.QueryRow(sqlStmt, nik)
	employee := &model.Employee{}
	err := row.Scan(
		&employee.ID,
		&employee.NIK,
		&employee.FirstName,
		&employee.LastName,
		&employee.Email,
	)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *EmployeeRepository) FetchEmployees() ([]model.Employee, error) {
	var sqlStmt string

	// Task:
	// buat query untuk mengambil data employees
	// simpan query ke variable sqlStmt
	// lihat model.Employee untuk field yang diambil dari database

	// TODO: answer here

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []model.Employee{}
	for rows.Next() {
		employee := model.Employee{}
		err := rows.Scan(
			&employee.ID,
			&employee.NIK,
			&employee.FirstName,
			&employee.LastName,
			&employee.Email,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
