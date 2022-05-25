package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

var _ = Describe("User", func() {
	BeforeEach(func() {
		// setup
		db, err := sql.Open("sqlite3", "./music-app.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
    		created_at DATETIME NOT NULL
		);`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO users (id, name, created_at) VALUES
			(1, 'John', '2020-01-01 00:00:00'),
			(2, 'Jane', '2020-01-01 00:00:00'),
			(3, 'Jack', '2020-01-01 00:00:00');`)

		if err != nil {
			panic(err)
		}
	})

	AfterEach(func() {
		db, err := sql.Open("sqlite3", "./music-app.db")
		if err != nil {
			panic(err)
		}
		_, err = db.Exec("DROP TABLE IF EXISTS users")
		if err != nil {
			panic(err)
		}
	})

	usersData := map[int]*model.User{
		1: &model.User{
			ID:        1,
			Name:      "John",
			CreatedAt: "2020-01-01T00:00:00Z",
		},
		2: &model.User{
			ID:        2,
			Name:      "Jane",
			CreatedAt: "2020-01-01T00:00:00Z",
		},
		3: &model.User{
			ID:        3,
			Name:      "Jack",
			CreatedAt: "2020-01-01T00:00:00Z",
		},
	}

	Describe("Fetch user by id", func() {
		It("should return user", func() {
			db, err := sql.Open("sqlite3", "./music-app.db")
			Expect(err).To(BeNil())

			userRepo := NewUserRepository(db)
			user, err := userRepo.FetchUserByID(1)
			Expect(err).To(BeNil())
			Expect(user).To(Equal(usersData[1]))
		})
	})

	Describe("Delete user by id", func() {
		It("should return user", func() {
			db, err := sql.Open("sqlite3", "./music-app.db")
			Expect(err).To(BeNil())

			userRepo := NewUserRepository(db)
			err = userRepo.DeleteUserByID(1)
			Expect(err).To(BeNil())

			user, err := userRepo.FetchUserByID(1)
			Expect(err).To(Equal(sql.ErrNoRows))
			Expect(user).To(BeNil())
		})
	})

})
