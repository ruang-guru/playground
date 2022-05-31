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

		_, err = db.Exec(`DROP TABLE unormal;`)
		if err != nil {
			panic(err)
		}

	})

	Describe("Check Migration", func() {
		It("should return number of column from unormal table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('unormal');`)
			Expect(err).To(BeNil())

			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).Should(BeNumerically(">=", 13))
			}
		})
	})

	Describe("Check Insert Latest data", func() {
		It("should return latest id", func() {
			dataExists, err := checkDataExists("00002")
			Expect(err).To(BeNil())
			Expect(dataExists).To(Equal(true))
		})
	})

})
