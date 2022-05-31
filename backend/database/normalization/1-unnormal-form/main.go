package main

// perhatikan gambar unormal-example.png
// identifikasikan-lah semua attribut yang ada, lalu buatlah table seperti pada gambar  dengan nama table "unormal"
// gambar digunakan untuk menampilkan bentuk tabel saja

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Unormal struct {
	IdMember         string `db:"id_member"`
	NamaMember       string `db:"nama_member"`
	TipeMember       string `db:"tipe_member"`
	KeteranganMember string `db:"keterangan_member"`
	IdTheater        string `db:"id_theater"`
	NamaTheater      string `db:"nama_theater"`
	Kota             string `db:"kota"`
	IdMovie          string `db:"id_movie"`
	NamaMovie        string `db:"nama_movie"`
	IdFaktur         string `db:"id_faktur"`
	WaktuTayang      string `db:"waktu_tayang"`
	HargaTiket       string `db:"harga_tiket"`
	QtyTicket        string `db:"qty_tiket"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS unormal (
		id_member VARCHAR(12),
		nama_member VARCHAR(12),
		tipe_member VARCHAR(10),
		keterangan_member VARCHAR(10),
		id_theater VARCHAR(10),
		nama_theater VARCHAR(10),
		kota VARCHAR(10),
		id_movie TEXT,
		nama_movie TEXT,
		id_faktur TEXT,
		waktu_tayang TEXT,
		harga_tiket TEXT,
		qty_tiket TEXT
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
			INSERT INTO 
			unormal (id_member, nama_member, tipe_member, keterangan_member, id_theater, nama_theater, kota, id_movie, nama_movie, id_faktur, waktu_tayang, harga_tiket, qty_tiket)
			VALUES
			("111", "Muri", "EPC", "Epic", "T01", "Paris Van Java", "Bandung", "M01, M03", "Orang kaya baru, Twice Land", "F001, F002", "22/01/2022 15:00, 22/01/2022 18:00", "30.000, 30.000", "1, 3"),
			("114", "Luga", "ELT", "Elite", "T02", "Grand Indonesia", "Jakarta", "M03, M04", "Twice Land, Escapee Room", "F003, F004", "22/02/2022 15:00, 22/02/2022 18:00", "30.000, 30.000", "2, 1")`)

	if err != nil {
		fmt.Println(err)
		fmt.Println("here")
	}

	return db, nil
}

func Rollback(db *sql.DB) {
	sqlStmt := `DROP TABLE unormal;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

}

// Jalankan main untuk melakukan migrasi database
func main() {
	db, err := Migrate()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM unormal")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var unormal Unormal
		err = rows.Scan(&unormal.IdMember, &unormal.NamaMember, &unormal.TipeMember, &unormal.KeteranganMember, &unormal.IdTheater, &unormal.NamaTheater, &unormal.Kota, &unormal.IdMovie, &unormal.NamaMovie, &unormal.IdFaktur, &unormal.WaktuTayang, &unormal.HargaTiket, &unormal.QtyTicket)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", unormal)
	}

	defer rows.Close()
}
