package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

var _ = Describe("SQL Inner Join", func() {
	BeforeEach(func() {
		// Setup database
		db, err := sql.Open("sqlite3", "./movie.db")
		if err != nil {
			panic(err)
		}

		// Create table movies
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS movies (
				id INTEGER PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				Description TEXT,
				year INTEGER NOT NULL,
			    director_id INTEGER NOT NULL,
			    genre_id INTEGER NOT NULL,
			    FOREIGN KEY(director_id) REFERENCES directors(id),
			    FOREIGN KEY(genre_id) REFERENCES genres(id)
			)
		`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO movies (title, Description, year, director_id, genre_id) VALUES 
			("The Shawshank Redemption", "Lorem Ipsum dolor sit amet", 1994, 1, 1),
			("The Godfather", "Lorem Ipsum dolor sit amet", 1972, 2, 2),
			("The Godfather: Part II", "Lorem Ipsum dolor sit amet", 1974, 2, 2),
			("The Dark Knight", "Lorem Ipsum dolor sit amet", 2008, 3, 3);`)
		if err != nil {
			panic(err)
		}

		// Create table directors
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS directors (
				id INTEGER PRIMARY KEY,
				name VARCHAR(255) NOT NULL
			)
		`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO directors (name) VALUES 
			("Frank Darabont"),
			("Francis Ford Coppola"),
			("Christopher Nolan");`)
		if err != nil {
			panic(err)
		}

		// Create table genres
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS genres (
				id INTEGER PRIMARY KEY,
				name VARCHAR(255) NOT NULL
			)
		`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO genres (name) VALUES 
			("Drama"),
			("Crime"),
			("Thriller");`)
		if err != nil {
			panic(err)
		}
	})
	AfterEach(func() {
		// Cleanup database
		db, err := sql.Open("sqlite3", "./movie.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`DROP TABLE movies`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`DROP TABLE directors`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`DROP TABLE genres`)
		if err != nil {
			panic(err)
		}

	})

	moviesData := []model.Movie{
		{
			ID:           1,
			Title:        "The Shawshank Redemption",
			Description:  "Lorem Ipsum dolor sit amet",
			Year:         1994,
			DirectorName: "Frank Darabont",
			GenreName:    "Drama",
		},
		{
			ID:           2,
			Title:        "The Godfather",
			Description:  "Lorem Ipsum dolor sit amet",
			Year:         1972,
			DirectorName: "Francis Ford Coppola",
			GenreName:    "Crime",
		},
		{
			ID:           3,
			Title:        "The Godfather: Part II",
			Description:  "Lorem Ipsum dolor sit amet",
			Year:         1974,
			DirectorName: "Francis Ford Coppola",
			GenreName:    "Crime",
		},
		{
			ID:           4,
			Title:        "The Dark Knight",
			Description:  "Lorem Ipsum dolor sit amet",
			Year:         2008,
			DirectorName: "Christopher Nolan",
			GenreName:    "Thriller",
		},
	}

	Describe("Fetch Movies by ID", func() {
		It("should return movie data", func() {
			db, err := sql.Open("sqlite3", "./movie.db")
			Expect(err).To(BeNil())

			movieRepo := NewMovieRepository(db)
			movie, err := movieRepo.FetchMovieByID(int64(1))
			Expect(err).To(BeNil())
			Expect(movie).To(Equal(&moviesData[0]))
		})
	})

	Describe("Fetch Movies", func() {
		It("should return all movie data", func() {
			db, err := sql.Open("sqlite3", "./movie.db")
			Expect(err).To(BeNil())

			movieRepo := NewMovieRepository(db)
			movies, err := movieRepo.FetchMovies()
			Expect(err).To(BeNil())
			Expect(movies).To(Equal(moviesData))
		})
	})
})
