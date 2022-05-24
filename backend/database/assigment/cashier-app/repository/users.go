package repository

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {
	return User{}, nil // TODO: replace this
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	return []User{}, nil // TODO: replace this
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	return nil, nil // TODO: replace this
}

func (u *UserRepository) InsertUser(username string, password string, role string, loggedin bool) error {
	return nil // TODO: replace this
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	return nil, nil // TODO: replace this
}
