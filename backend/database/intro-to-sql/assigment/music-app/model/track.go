package model

type Track struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Artist    string `db:"artist"`
	CreatedAt string `db:"created_at"`
}
