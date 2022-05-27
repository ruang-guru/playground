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
		_, err := sql.Open("sqlite3", "./surrogate.db")
		if err != nil {
			panic(err)
		}
		_, err = Migrate()
		if err != nil {
			panic(err)
		}

	})

	AfterEach(func() {
		db, err := sql.Open("sqlite3", "./surrogate.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE school_a_cp;
		DROP TABLE school_b_cp;
		DROP TABLE surrogate_table_cp;`)

		if err != nil {
			panic(err)
		}

		//comment this line below if you want to see the table, before running test case
		os.Remove("./surrogate.db")
	})

	Describe("Check Migration 1", func() {
		It("should return number of column from school_a_cp table", func() {
			db, err := sql.Open("sqlite3", "./surrogate.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('school_a_cp');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(5))
			}
		})
	})

	Describe("Check Migration 2", func() {
		It("should return number of column from school_b_cp table", func() {
			db, err := sql.Open("sqlite3", "./surrogate.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('school_b_cp');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(3))
			}
		})
	})

	Describe("Check Migration 3", func() {
		It("should return number of column from surrogate_table_cp table", func() {
			db, err := sql.Open("sqlite3", "./surrogate.db")
			Expect(err).To(BeNil())
			CheckMigration, err := db.Query(`SELECT COUNT(*) FROM pragma_table_info('surrogate_table_cp');`)
			for CheckMigration.Next() {
				var success int
				err = CheckMigration.Scan(&success)
				Expect(err).To(BeNil())
				Expect(success).To(Equal(6))
			}
		})
	})

})
