package model

type User struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
}

type UserPlaylist struct {
	UserID       int64  `db:"user_id"`
	UserName     string `db:"user_name"`
	PlaylistID   int64  `db:"playlist_id"`
	PlaylistName string `db:"playlist_name"`
}
