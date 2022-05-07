package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) LoadOrCreate() ([]Product, error) {
	return []Product{}, nil // TODO: replace this
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	return []Product{}, nil // TODO: replace this
}
