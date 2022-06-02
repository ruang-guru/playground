package main

// pada contoh gambar surrogate_key, diketahui terdapat 2 table pendaftaran sekolah, dimana kedua table tersebut
// memilik kolom regirstration_no, name, dan percentage.
// gabunglah kedua table tersebut, dan buatlah surrogate key pada table yang baru.

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SchoolA struct {
	RegistrationNo string  `db:"registration_no"`
	Name           string  `db:"name"`
	Percentage     float64 `db:"percentage"`
}

type SchoolB struct {
	RegistrationNo string  `db:"registration_no"`
	Name           string  `db:"name"`
	Percentage     float64 `db:"percentage"`
}

type SurrogateTable struct {
	Id             int     `db:"id"`
	RegistrationNo string  `db:"registration_no"`
	Name           string  `db:"name"`
	Percentage     float64 `db:"percentage"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./surrogate.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS school_a (
		registration_no Varchar(16) PRIMARY KEY,
		name TEXT,
		percentage REAL
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			school_a (registration_no, name, percentage)
			VALUES 
			    ("1000", "SMA Negeri 1", 0.5),
				("2000", "SMA Negeri 2", 0.5),
				("3000", "SMA Negeri 3", 0.5),
				("4000", "SMA Negeri 4", 0.5),
				("5000", "SMA Negeri 5", 0.5);`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS school_b (
		registration_no Varchar(16) PRIMARY KEY,
		name TEXT,
		percentage REAL
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			school_b (registration_no, name, percentage)
			VALUES 
			    ("1000", "SMA Negeri 6", 0.5),
				("2000", "SMA Negeri 7", 0.5),
				("3000", "SMA Negeri 8", 0.5),
				("4000", "SMA Negeri 9", 0.5),
				("5000", "SMA Negeri 10", 0.5);`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS surrogate_table (
		id INTEGER PRIMARY KEY,
		registration_no Varchar(16),
		name TEXT,
		percentage REAL
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			surrogate_table (registration_no, name, percentage)
			VALUES 
				("1000", "SMA Negeri 1", 0.5),
				("2000", "SMA Negeri 2", 0.5),
				("3000", "SMA Negeri 3", 0.5),
				("4000", "SMA Negeri 4", 0.5),
				("5000", "SMA Negeri 5", 0.5),
			    ("1000", "SMA Negeri 6", 0.5),
				("2000", "SMA Negeri 7", 0.5),
				("3000", "SMA Negeri 8", 0.5),
				("4000", "SMA Negeri 9", 0.5),
				("5000", "SMA Negeri 10", 0.5);`)
	// Kenapa tidak perlu memasukkan id?
	// karena PRIMARY KEY pada SQLite adalah alias untuk ROWID
	// ROWID sendiri adalah 64-bit signed integer yang digunakan untuk
	// mengidentifikasi baris pada table secara unik
	// secara default tiap table pada SQLite memiliki ROWID

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Jalankan main untuk melakukan migrasi database
func main() {
	db, err := Migrate()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM school_a")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var schoolA SchoolA
		err = rows.Scan(&schoolA.RegistrationNo, &schoolA.Name, &schoolA.Percentage)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", schoolA)
	}

	rows, err = db.Query("SELECT * FROM school_b")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var schoolB SchoolB
		err = rows.Scan(&schoolB.RegistrationNo, &schoolB.Name, &schoolB.Percentage)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", schoolB)
	}

	rows, err = db.Query("SELECT * FROM surrogate_table")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var surrogateTable SurrogateTable
		err = rows.Scan(&surrogateTable.Id, &surrogateTable.RegistrationNo, &surrogateTable.Name, &surrogateTable.Percentage)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", surrogateTable)
	}

	defer rows.Close()
}
