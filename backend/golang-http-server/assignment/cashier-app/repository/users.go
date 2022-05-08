package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	return []User{}, nil // TODO: replace this
}

func (u *UserRepository) SelectAll() ([]User, error) {
	return []User{}, nil // TODO: replace this
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	return nil, nil // TODO: replace this
}

func (u *UserRepository) Save(users []User) error {
	return nil // TODO: replace this
}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	// TODO: answer here
}
