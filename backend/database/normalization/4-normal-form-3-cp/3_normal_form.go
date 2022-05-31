package main

// Normalisasi database dalam bentuk 3NF bertujuan untuk menghilangkan seluruh atribut atau field yang tidak berhubungan dengan primary key.
// Dengan demikian tidak ada ketergantungan transitif pada setiap kandidat key
// fungsi normalisasi 3NF antara lain :
// 1. Memenuhi semua persyaratan dari bentuk normal kedua.
// 2. Menghapus kolom yang tidak tergantung pada primary key.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Rekap struct {
	NoBon     int
	Discount  int
	Total     int
	Bayar     int
	Kembalian int
	KodeKasir string
	Tanggal   string
	Waktu     string
}

type DetailRekap struct {
	NoBon      int
	KodeBarang string
	Harga      int
	Jumlah     int
}

type Barang struct {
	KodeBarang string
	NamaBarang string
	Harga      int
}

type Kasir struct {
	KodeKasir string
	NamaKasir string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
// Buatlah tabel dengan nama rekap, rekap_detail, barang, dan kasir
// Lalu insert data ke masing-masing tabel seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE rekap ... ` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO rekap ... `) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE rekap_detail ... ` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO rekap_detail ... `) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE barang ... ` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSEERT INTO barang ... `) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE kasir ... ` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO kasir ... `) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkNoBonExists:
// checkNoBonExists digunakan untuk menghitung jumlah data yang ada berdasarkan no_bon
func checkNoBonExists(noBon string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi countRekapDetailByNoBon:
// countRekapDetailByNoBon digunakan untuk menghitung jumlah rekap detail yang ada berdasarkan no_bon
func countRekapDetailByNoBon(noBon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return 0, err
	}
	return countBon, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkBarangExists:
func checkBarangExists(kodeBarang string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `...` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeBarang)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkKasirExists:
func checkKasirExists(kodeKasir string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `...` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeKasir)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}
