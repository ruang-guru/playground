package main

type EmployeeRow struct {
	ID        int
	Name      string
	Position  string
	Salary    int
	ManagerID int
}
type EmployeeDB []EmployeeRow

func NewEmployeeDB() *EmployeeDB {
	return &EmployeeDB{}
}

func (db *EmployeeDB) Where(id int) *EmployeeRow {
	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			return &(*db)[i]
		}
	}
	return nil
}

func (db *EmployeeDB) Insert(name string, position string, salary int, managerID int) {
	(*db) = append(*db, EmployeeRow{
		ID:        len(*db) + 1,
		Name:      name,
		Position:  position,
		Salary:    salary,
		ManagerID: managerID,
	})
}

func (db *EmployeeDB) Update(id int, name string, position string, salary int, managerID int) {
	// TODO: answer here
}

func (db *EmployeeDB) Delete(id int) {
	// TODO: answer here
}
