package main

// pada tahap normalisai 1 (1NF), kita akan menyederhanakan bentuk unormal sesuai dengan kaidah bentuk normalisasi 1
// - Setiap kolom hanya memiliki satu nilai (atomic value)
// - Tipe data yang disimpan pada suatu kolom harus sama
// - Nama kolom harus unik
// - Urutan penyimpanan data tidak perlu diperhatikan
// untuk id_member pada gambar seharusnya sama tidak berubah
// nilai id_theater pada gambar kurang sesuai, ikuti yang ada pada program
// gambar digunakan untuk menampilkan bentuk tabel saja

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Faktur struct {
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
	HargaTiket       int    `db:"harga_tiket"`
	QtyTicket        int    `db:"qty_tiket"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS faktur_1nf (
		id_member VARCHAR(12) PRIMARY KEY,
		nama_member VARCHAR(12),
		tipe_member VARCHAR(10),
		keterangan_member VARCHAR(10),
		id_theater VARCHAR(10),
		nama_theater VARCHAR(10),
		kota VARCHAR(10),
		id_movie VARCHAR(12),
		nama_movie VARCHAR(12),
		id_faktur VARCHAR(12),
		waktu_tayang VARCHAR(12),
		harga_tiket INTEGER,
		qty_tiket INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
	INSERT INTO 
	faktur_1nf (id_member, nama_member, tipe_member, keterangan_member, id_theater, nama_theater, kota, id_movie, nama_movie, id_faktur, waktu_tayang, harga_tiket, qty_tiket)
	VALUES
	("111", "Muri", "EPC", "Epic", "T01", "Paris Van Java", "Bandung", "M01", "Orang kaya baru", "F001", "22/01/2022 15:00", 30000, 1),
	("114", "Luga", "ELT", "Elite", "T02", "Grand Indonesia", "Jakarta", "M03", "Twice Land", "F003", "22/02/2022 15:00", 30000, 2)`)
	// sekarang tiap kolom tidak ada yang memiliki multi value

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
		err = rows.Scan(&faktur.IdMember, &faktur.NamaMember, &faktur.TipeMember, &faktur.KeteranganMember, &faktur.IdTheater, &faktur.NamaTheater, &faktur.Kota, &faktur.IdMovie, &faktur.NamaMovie, &faktur.IdFaktur, &faktur.WaktuTayang, &faktur.HargaTiket, &faktur.QtyTicket)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", faktur)
	}

	defer rows.Close()
}
