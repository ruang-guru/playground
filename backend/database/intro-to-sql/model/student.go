package model

type Student struct {
	ID   int    `db:"id"`
	NIM  string `db:"nim"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
