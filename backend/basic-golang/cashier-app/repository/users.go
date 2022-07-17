package repository

import (
	"fmt"
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
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

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	return nil, nil // TODO: replace this
}

func (u *UserRepository) Logout(username string) error {
	return nil // TODO: replace this
}

func (u *UserRepository) Save(users []User) error {
	return nil // TODO: replace this
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (u *UserRepository) LogoutAll() error {
	return nil // TODO: replace this
}
