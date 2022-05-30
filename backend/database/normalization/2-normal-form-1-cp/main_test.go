package main

import (
	"database/sql"

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

		_, err = db.Exec(`DROP TABLE rekap;`)
		if err != nil {
			panic(err)
		}

	})

	Describe("Check Migration", func() {
		It("should return number of column from rekap table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('rekap');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(13))
			}
		})
	})

	Describe("Check Insert Latest data", func() {
		It("should return latest id", func() {
			countData, err := countByNoBon("00001")
			Expect(err).To(BeNil())
			Expect(countData).To(Equal(4))

			countData, err = countByNoBon("00002")
			Expect(err).To(BeNil())
			Expect(countData).To(Equal(3))
		})
	})

})
