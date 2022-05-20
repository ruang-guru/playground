package repository

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

type UserRepositoryInterface interface {
	FetchUserByID(id int) (*model.User, error)
	DeleteUserByID(id int) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (ur *UserRepository) FetchUserByID(id int) (*model.User, error) {
	// var sqlStatement string

	// Task 1: lengkapi statement SQL untuk mengambil data user berdasarkan id
	// TODO: answer here
	sqlStatement := "SELECT DISTINCT id, name, created_at FROM users WHERE id=?"

	var user model.User
	// Task 2: buatlah query dengan prepared statement dengan statement SQL yang sudah di lengkapi	diatas
	// TODO: answer here
	res := ur.db.QueryRow(sqlStatement, id)

	err := res.Scan(&user.ID, &user.Name, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) DeleteUserByID(id int) error {
	// var sqlStatement string

	// Task 1: lengkapi statement SQL untuk menghapus data user berdasarkan id
	// TODO: answer here
	sqlStatement := "DELETE FROM users WHERE id=?"

	// Task 2: buatlah exec query dengan prepared statement dengan statement SQL yang sudah di lengkapi	diatas
	// TODO: answer here
	_, err := ur.db.Exec(sqlStatement, id)
	return err
}
