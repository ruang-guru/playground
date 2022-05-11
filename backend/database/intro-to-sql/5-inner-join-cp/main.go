package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db}
}

func (r *MovieRepository) FetchMovies() ([]model.Movie, error) {
	var sqlStmt string

	// Task : create a query to fetch all movies
	// 1. create a query to fetch all movies
	// 2. use inner join to fetch all movies and their genres and their directors

	// TODO: answer here

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []model.Movie
	for rows.Next() {
		var m model.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Description,
			&m.Year,
			&m.GenreName,
			&m.DirectorName,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func (r *MovieRepository) FetchMovieByID(id int64) (*model.Movie, error) {
	var sqlStmt string

	// Task : create a query to fetch a movie by id
	// 1. create a query to fetch a movie by id
	// 2. use inner join to fetch all movies and their genres and their directors

	// TODO: answer here

	row := r.db.QueryRow(sqlStmt, id)

	var m model.Movie
	err := row.Scan(
		&m.ID,
		&m.Title,
		&m.Description,
		&m.Year,
		&m.GenreName,
		&m.DirectorName,
	)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
