package repository

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

type PlaylistRepositoryInterface interface {
	FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error)
	CreatePlaylist(playlist model.Playlist) error
	UpdateUserPlaylist(userID int64, playlist model.Playlist) error
	FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error)
}

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) PlaylistRepositoryInterface {
	return &PlaylistRepository{db}
}

func (p *PlaylistRepository) FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error) {
	var sqlStatement string
	var playlists []model.UserPlaylist

	// Task 1: Buat query untuk mengambil playlist yang dimiliki user
	// TODO: answer here
	sqlStatement = "SELECT DISTINCT users.id, users.name, playlists.id, playlists.name FROM users JOIN playlists ON users.id = playlists.user_id WHERE users.id = ?"

	// Task 2: Buat execute statement untuk mengambil playlist yang dimiliki user
	// TODO: answer here
	res, err := p.db.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	// Task 3: Buat looping untuk mengambil data playlist yang dimiliki user
	// TODO: answer here
	for res.Next() {
		playlist := model.UserPlaylist{}
		err := res.Scan(
			&playlist.UserID,
			&playlist.UserName,
			&playlist.PlaylistID,
			&playlist.PlaylistName,
		)

		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (p *PlaylistRepository) CreatePlaylist(playlist model.Playlist) error {

	// Task 1 : Buat query untuk menambahkan playlist baru
	// TODO: answer here
	sqlStatement := "INSERT INTO playlists (name, user_id, created_at) values (?,?,?)"

	// Task 2 := Buat execute statement untuk menambahkan playlist baru
	// TODO: answer here
	_, err := p.db.Exec(sqlStatement, playlist.Name, playlist.UserID, playlist.CreatedAt)

	return err
}

func (p *PlaylistRepository) UpdateUserPlaylist(userID int64, playlist model.Playlist) error {

	// Task 1 : Buat query untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here
	sqlStatement := "UPDATE playlists SET name=?, user_id=?, created_at=? WHERE user_id=?"

	// Task 2 : Buat execute statement untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here
	_, err := p.db.Exec(sqlStatement, playlist.Name, playlist.UserID, playlist.CreatedAt, userID)

	return err
}

func (p *PlaylistRepository) FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error) {

	// Task 1 : Buat query untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	sqlStatement := "SELECT DISTINCT playlists.id, playlists.name, tracks.id, tracks.name, tracks.artist FROM playlist_tracks JOIN tracks ON playlist_tracks.track_id = tracks.id JOIN playlists ON playlist_tracks.playlist_id = playlists.id WHERE playlists.id = ?"

	// Task 2 : Buat execute statement untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	res, err := p.db.Query(sqlStatement, playlistID)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var playlistTracks []model.PlaylistTrack
	// Task 3 : Buat looping untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here
	for res.Next() {
		track := model.PlaylistTrack{}
		res.Scan(&track.PlaylistID, &track.PlaylistName, &track.TrackID, &track.TrackName, &track.TrackArtist)
		playlistTracks = append(playlistTracks, track)
	}

	return playlistTracks, nil
}
