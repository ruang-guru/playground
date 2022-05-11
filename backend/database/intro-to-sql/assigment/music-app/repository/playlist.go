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

	// Task 2: Buat execute statement untuk mengambil playlist yang dimiliki user
	// TODO: answer here

	// Task 3: Buat looping untuk mengambil data playlist yang dimiliki user
	var playlist model.UserPlaylist
	// TODO: answer here

	return playlists, nil
}

func (p *PlaylistRepository) CreatePlaylist(playlist model.Playlist) error {
	var sqlStatement string

	// Task 1 : Buat query untuk menambahkan playlist baru
	// TODO: answer here

	// Task 2 := Buat execute statement untuk menambahkan playlist baru
	// TODO: answer here

	return nil
}

func (p *PlaylistRepository) UpdateUserPlaylist(userID int64, playlist model.Playlist) error {
	var sqlStatement string

	// Task 1 : Buat query untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here

	// Task 2 : Buat execute statement untuk mengupdate playlist name dengan id playlist tertentu yang dimiliki user tertentu
	// TODO: answer here

	return nil
}

func (p *PlaylistRepository) FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error) {
	var sqlStatement string

	// Task 1 : Buat query untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here

	// Task 2 : Buat execute statement untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here

	var playlistTracks []model.PlaylistTrack
	// Task 3 : Buat looping untuk mengambil track yang dimiliki playlist tertentu
	// TODO: answer here

	return playlistTracks, nil
}

