package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SQL Create", func() {
	BeforeEach(func() {
		_, err := sql.Open("sqlite3", "./normalize-cp.db")
		if err != nil {
			panic(err)
		}
		_, err = Migrate()
		if err != nil {
			panic(err)
		}

	})

	AfterEach(func() {
		db, err := sql.Open("sqlite3", "./normalize-cp.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE rekap_2nf;
		DROP TABLE barang_2nf;
		DROP TABLE kasir_2nf;`)

		if err != nil {
			panic(err)
		}

		//comment this line below if you want to see the table, before running test case
		os.Remove("./normalize-cp.db")
	})

	Describe("Check Migration", func() {
		It("should return number of column from rekap_2nf table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('rekap_2nf');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(13))
			}
		})
	})

	Describe("Check Migration", func() {
		It("should return number of column from barang_2nf table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('barang_2nf');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(3))
			}
		})
	})

	Describe("Check Migration", func() {
		It("should return number of column from kasir_2nf table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('kasir_2nf');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(2))
			}
		})
	})

	Describe("Check Insert Latest data 1", func() {
		It("should return latest no_bon", func() {
			lastInsertData, err := checkLatestNoBon("00002")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))

		})
	})

	Describe("Check Insert Latest data 2", func() {
		It("should return latest no_barang", func() {
			lastInsertData, err := checkLatestNoBarang("B005")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))

		})
	})

	Describe("Check Insert Latest data 3", func() {
		It("should return latest no_kasir", func() {
			lastInsertData, err := checkLatestNoKasir("K02")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))

		})
	})

})
