package repository

import (
	"errors"
	"sync"

	"github.com/ruang-guru/playground/backend/performance-testing/server/entity"
)

type repository struct {
	movies map[int]entity.Movie
}

type Repository interface {
	Get(id int) (interface{}, error)
	Store(movie *entity.Movie) error
	GetAll() (map[int]entity.Movie, error)
}

func NewRepo() Repository {
	return &repository{
		movies: make(map[int]entity.Movie),
	}
}

var mtx sync.Mutex

func (r *repository) GetAll() (map[int]entity.Movie, error) {
	if len(r.movies) > 0 {
		return r.movies, nil
	}
	return nil, errors.New("data not found")
}

func (r *repository) Get(id int) (interface{}, error) {
	if v, ok := r.movies[id]; ok {
		return v, nil
	}
	return nil, errors.New("data not found")

}

func (r *repository) Store(movie *entity.Movie) error {
	id := len(r.movies) + 1
	movie.ID = id
	mtx.Lock()
	r.movies[id] = *movie
	mtx.Unlock()
	return nil
}
