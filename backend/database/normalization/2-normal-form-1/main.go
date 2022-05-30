package main

// pada tahap normalisai 1 (1NF), kita akan menyederhanakan bentuk unormal sesuai dengan kaidah bentuk normalisasi 1
// dengan menghilangkan duplikasi kolom dari tabel yang sama, buat tabel terpisah untuk masing - masing kelompok data terkait, dan identifikasikan
// setiap baris dengan kolom yang unik (primary key)

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Faktur struct {
	Id          int    `db:"id"`
	NoFaktur    int    `db:"no_faktur"`
	Tanggal     string `db:"tanggal"`
	NoPolisi    string `db:"no_polisi"`
	Warna       string `db:"warna"`
	Merek       string `db:"merek"`
	Tahun       string `db:"tahun"`
	MekanikId   string `db:"mekanik_id"`
	NamaMekanik string `db:"nama_mekanik"`
	KodeParts   string `db:"kode_parts"`
}

type BonPembelian struct {
	Id        int    `db:"id"`
	NamaParts string `db:"nama_parts"`
	Kuantitas int    `db:"kuantitas"`
	Harga     int    `db:"harga"`
	Discount  int    `db:"discount"`
	Jumlah    int    `db:"jumlah"`
	Potongan  int    `db:"potongan"`
	Total     int    `db:"total"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS faktur_1nf (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		no_faktur INTEGER,
		tanggal VARCHAR(12),
		no_polisi VARCHAR(10),
		warna VARCHAR(10),
		merek VARCHAR(10),
		tahun VARCHAR(10),
		mekanik_id VARCHAR(10),
		nama_mekanik VARCHAR(10),
		kode_parts VARCHAR(10)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
			INSERT INTO 
			faktur_1nf (no_faktur, tanggal, no_polisi, warna, merek, tahun, mekanik_id, nama_mekanik, kode_parts)
			VALUES 
			    (05103214, "2020-01-01", "B3117lB", "Biru", "Supra X", "2020", "DDE", "Djoko Dewanto", "20W501000CC"),
				(05103214, "2020-01-01", "B3117lB", "Biru", "Supra X", "2020", "DDE", "Djoko Dewanto", "SERV001");`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS bon_pembelian_1nf (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nama_parts VARCHAR(10),
		kuantitas INTEGER,
		harga INTEGER,
		discount INTEGER,
		jumlah INTEGER,
		potongan INTEGER,
		total INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			bon_pembelian_1nf (nama_parts, kuantitas, harga, discount, jumlah, potongan, total)
			VALUES 
			    ("Oli Top 1000cc", 2, 27000, 1000, 52000, 2000, 73000),
				("Engine Tune Up", 1, 25000, 2000, 23000, 2000, 73000);`)

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

	rows, err := db.Query("SELECT * FROM faktur_1nf")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var faktur Faktur
		err = rows.Scan(&faktur.Id, &faktur.NoFaktur, &faktur.Tanggal, &faktur.NoPolisi, &faktur.Warna, &faktur.Merek, &faktur.Tahun, &faktur.MekanikId, &faktur.NamaMekanik, &faktur.KodeParts)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", faktur)
	}

	rows, err = db.Query("SELECT * FROM bon_pembelian_1nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var bon_pembelian_1nf BonPembelian
		err = rows.Scan(&bon_pembelian_1nf.Id, &bon_pembelian_1nf.NamaParts, &bon_pembelian_1nf.Kuantitas, &bon_pembelian_1nf.Harga, &bon_pembelian_1nf.Discount, &bon_pembelian_1nf.Jumlah, &bon_pembelian_1nf.Potongan, &bon_pembelian_1nf.Total)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", bon_pembelian_1nf)
	}

	defer rows.Close()
}
