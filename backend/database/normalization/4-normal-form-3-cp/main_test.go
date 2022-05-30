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
		DROP TABLE rekap_3nf;
		DROP TABLE rekap_detail_3nf;
		DROP TABLE barang_3nf;
		DROP TABLE kasir_3nf;`)

		if err != nil {
			panic(err)
		}

		//comment this line below if you want to see the table, before running test case
		os.Remove("./normalize-cp.db")
	})

	Describe("Check Migration 1", func() {
		It("should return number of column from rekap_3nf table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('rekap_3nf');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(8))
			}
		})
	})

	Describe("Check Migration 2", func() {
		It("should return number of column from rekap_detail_3nf table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('rekap_detail_3nf');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(4))
			}
		})
	})

	Describe("Check Migration 3", func() {
		It("should return number of column from barang_3nf table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('barang_3nf');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(3))
			}
		})
	})

	Describe("Check Migration 4", func() {
		It("should return number of column from kasir_3nf table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('kasir_3nf');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(2))
			}
		})
	})

	Describe("Check Insert Latest no_bon in rekap table", func() {
		It("should return latest no_bon", func() {
			lastInsertData, err := checkLatestNoBon("00002")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))

		})
	})

	Describe("Check Insert Latest no_bon in rekap_detail table", func() {
		It("should return latest no_bon", func() {
			lastInsertData, err := checkLatestNoBonDetail("00002")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))

		})
	})

	Describe("Check Insert Latest no_barang in barang table", func() {
		It("should return latest no_barang", func() {
			lastInsertData, err := checkLatestNoBarang("B005")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))

		})
	})

	Describe("Check Insert Latest no_kasir in kasir table", func() {
		It("should return latest no_kasir", func() {
			lastInsertData, err := checkLatestNoKasir("K02")
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))

		})
	})

})
