package model

type Playlist struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	UserID    int64  `db:"user_id"`
	CreatedAt string `db:"created_at"`
}

type PlaylistTrack struct {
	PlaylistID   int64  `db:"playlist_id"`
	PlaylistName string `db:"playlist_name"`
	TrackID      int64  `db:"track_id"`
	TrackName    string `db:"track_name"`
	TrackArtist  string `db:"track_artist"`
}
