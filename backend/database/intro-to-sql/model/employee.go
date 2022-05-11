package model

type Employee struct {
	ID        int    `db:"id"`
	NIK       string `db:"nik"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}
