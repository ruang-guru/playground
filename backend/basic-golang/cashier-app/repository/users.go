package repository

import (
	"errors"
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
	records, err := u.db.Load("users")
	if err != nil {
		records = [][]string{
			{"username", "password", "loggedin"},
		}
		if err := u.db.Save("users", records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		isLoggedIn, _ := strconv.ParseBool(records[i][2])

		userItem := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: isLoggedIn,
		}
		result = append(result, userItem)
	}
	return result, nil // TODO: replace this
}

func (u *UserRepository) SelectAll() ([]User, error) {
	return u.LoadOrCreate() // TODO: replace this
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	var loggedUser *string
	for _, val := range users {

		err := u.changeStatus(val.Username, false)
		if err != nil {
			return nil, err
		}

	}

	for _, val := range users {
		if val.Username == username && val.Password == password {
			err := u.changeStatus(username, true)
			if err != nil {
				return nil, err
			}
			loggedUser = &val.Username
			return loggedUser, nil
		}
	}
	err = errors.New("Login Failed")
	return nil, err // TODO: replace this
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, val := range users {
		if val.Loggedin {
			return &val.Username, nil
		}
	}
	err = errors.New("no user is logged in")
	return nil, err // TODO: replace this
}

func (u *UserRepository) Logout(username string) error {
	loggedUser, err := u.FindLoggedinUser()
	if err != nil {
		return err
	}

	if *loggedUser == username {
		return u.changeStatus(username, false)
	}

	err = errors.New("rejects the logout")
	return err // TODO: replace this
}

func (u *UserRepository) Save(users []User) error {
	records := [][]string{
		{
			"username", "password", "loggedin",
		},
	}

	for _, val := range users {
		records = append(records, []string{
			val.Username, val.Password, strconv.FormatBool(val.Loggedin),
		})
	}
	u.db.Delete("users")
	return u.db.Save("users", records) // TODO: replace this
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for i, val := range users {
		if val.Username == username {
			users[i].Loggedin = status
			break
		}
	}

	return u.Save(users) // TODO: replace this
}

func (u *UserRepository) LogoutAll() error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for i, _ := range users {
		users[i].Loggedin = false
	}

	return u.Save(users) // TODO: replace this
}
