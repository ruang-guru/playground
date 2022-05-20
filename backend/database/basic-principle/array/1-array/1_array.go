package main

import "fmt"

type StudentRow struct {
	ID          int
	NIM         string
	StudentName string
	Class       string
	SPP         int
	NoHP        string
}
type StudentDB []StudentRow

func main() {
	var students StudentDB
	students.Insert("John", "123456781", "B-01", 100000, "08123456781")
	students.Insert("Jane", "123456782", "C-04", 300000, "08123456782")
	students.Insert("Gina", "123456783", "A-03", 230000, "08123456783")
	students.Insert("Vina", "123456784", "C-01", 240000, "08123456784")

	fmt.Println(students.Where(3))
}

func (db *StudentDB) Insert(name string, nim string, class string, spp int, nohp string) {
	(*db) = append(*db, StudentRow{
		ID:          len(*db) + 1,
		NIM:         nim,
		StudentName: name,
		Class:       class,
		SPP:         spp,
		NoHP:        nohp,
	})
}

func (db *StudentDB) Where(id int) *StudentRow {
	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			return &(*db)[i]
		}
	}
	return nil
}
