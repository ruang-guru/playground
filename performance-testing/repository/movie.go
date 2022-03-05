package repository

import (
	"errors"
	"sync"
)

type repository struct {
	movies map[int]Movie
}

type Repository interface {
	Get(id int) (interface{}, error)
	Store(movie *Movie) error
	GetAll() (map[int]Movie, error)
}

type Movie struct {
	Id      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

func NewRepo() Repository {
	return &repository{
		movies: make(map[int]Movie),
	}
}

var mtx sync.Mutex

func (r *repository) GetAll() (map[int]Movie, error) {
	if len(r.movies) > 0 {
		return r.movies, nil
	} else {
		return nil, errors.New("data not found")
	}
}

func (r *repository) Get(id int) (interface{}, error) {
	if v, ok := r.movies[id]; ok {
		return v, nil
	} else {
		return nil, errors.New("data not found")
	}
}

func (r *repository) Store(movie *Movie) error {
	id := len(r.movies) + 1
	movie.Id = id
	mtx.Lock()
	r.movies[id] = *movie
	mtx.Unlock()
	return nil
}
