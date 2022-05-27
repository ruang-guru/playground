package main

// Syarat untuk menerapkan normalisasi bentuk kedua ini adalah data telah dibentuk dalam 1NF
// fungsi normalisasi 2NF antara lain :
// 1. Menghapus beberapa subset data yang ada pada tabel dan menempatkan mereka pada tabel terpisah.
// 2. Menciptakan hubungan antara tabel baru dan tabel lama dengan menciptakan foreign key.
// 3. Tidak ada atribut dalam tabel yang secara fungsional bergantung pada candidate key tabel tersebut.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Rekap struct {
	NoBon      string
	KodeBarang string
	Harga      int
	Jumlah     int
	Biaya      int
	SubTotal   int
	Discount   int
	Total      int
	Bayar      int
	Kembalian  int
	KodeKasir  string
	Tanggal    string
	Waktu      string
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
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}
	sqlStmt := `CREATE TABLE ... ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`INSERT INTO ... VALUES ... ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE ... ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO ... VALUES ... ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE ... ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO ... VALUES ... ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestNoBon:
func checkLatestNoBon(no_bon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this

	row := db.QueryRow(sqlStmt, no_bon)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestNoBarang:
func checkLatestNoBarang(kode_barang string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `...` // TODO: replace this

	row := db.QueryRow(sqlStmt, kode_barang)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestNoKasir:
func checkLatestNoKasir(kode_kasir string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `...` // TODO: replace this

	row := db.QueryRow(sqlStmt, kode_kasir)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}

//insert value table rekap_2nf
// ("00001", "B001", 4500, 3, 13500, 13500, 0, 13500, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
// ("00001", "B002", 22500, 1, 22500, 36000, 0, 36000, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
// ("00001", "B003", 1500, 4, 6000, 42000, 0, 42000, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
// ("00001", "B004", 17500, 2, 35000, 77000, 0, 77000, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
// ("00002", "B001", 4500, 1, 4500, 4500, 0, 4500, 17500, 0, "K02", "04-05-2022", "12:00:00"),
// ("00002", "B004", 17400, 1, 17500, 22000, 0, 22000, 117500, 0, "K02", "04-05-2022", "12:00:00"),
// ("00002", "BOO5", 100000, 1, 100000, 117500, 0, 117500, 117500, 0, "K02", "04-05-2022", "12:00:00")

//insert value table barang_2nf
// ("B001", "Disket", 4500),
// ("B002", "Refil Tinta", 22500),
// ("B003", "CD Blank", 1500),
// ("B004", "Mouse", 17500),
// ("B005", "Flash Disk", 100000)

//insert value table kasir_2nf
// ("K01", "Rosi"),
// ("K02", "Dewi")
