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

		_, err = db.Exec(`DROP TABLE unormal;`)
		if err != nil {
			panic(err)
		}

		//comment this line below if you want to see the table, before running test case
		os.Remove("./normalize-cp.db")
	})

	Describe("Check Migration", func() {
		It("should return number of column from unormal table", func() {
			db, err := sql.Open("sqlite3", "./normalize-cp.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('unormal');`)
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
			lastInsertData, err := checkLatestId(2)
			Expect(err).To(BeNil())
			Expect(lastInsertData).To(Equal(1))
		})
	})

})
