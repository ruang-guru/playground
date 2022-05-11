package repository

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

var _ = Describe("Playlist", func() {
	BeforeEach(func() {
		db, err := sql.Open("sqlite3", "./music-app.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
    		created_at DATETIME NOT NULL
		);`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO users (id, name, created_at) VALUES
			(1, 'John', '2020-01-01 00:00:00');`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS playlists (
    		id INTEGER PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
    		created_at DATETIME NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id));`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO playlists (id, name, created_at, user_id) VALUES
		(1, 'Chill Music', '2020-01-01 00:00:00', 1),
		(2, 'POP 2022', '2020-01-01 00:00:00', 1),
		(3, 'Rock', '2020-01-01 00:00:00', 1);`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tracks (
			id INTEGER PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
    		artist VARCHAR(255) NOT NULL,
			created_at DATETIME NOT NULL);`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO tracks (id, name, artist, created_at) VALUES
		(1, 'Music A', 'Robert', '2020-01-01 00:00:00'),
		(2, 'Music B', 'Garnier', '2020-01-01 00:00:00'),
    	(3, 'Music C', 'Soppy', '2020-01-01 00:00:00'),
    	(4, 'Music D', 'Nicky', '2020-01-01 00:00:00'),
    	(5, 'Music E', 'Rony', '2020-01-01 00:00:00'),
    	(6, 'Music F', 'Jane', '2020-01-01 00:00:00');`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS playlist_tracks (
			playlist_id INTEGER NOT NULL,
			track_id INTEGER NOT NULL,
			FOREIGN KEY (playlist_id) REFERENCES playlists(id),
			FOREIGN KEY (track_id) REFERENCES tracks(id));`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`INSERT INTO playlist_tracks (playlist_id, track_id) VALUES 
		(1, 1),
		(1, 2),
		(1, 3),
    	(2, 4),
		(2, 5),
		(3, 6);`)
		if err != nil {
			panic(err)
		}
	})

	AfterEach(func() {
		os.Remove("./music-app.db")
	})

	usersPlaylist := []model.UserPlaylist{
		{
			UserID:       1,
			UserName:     "John",
			PlaylistID:   1,
			PlaylistName: "Chill Music",
		},
		{
			UserID:       1,
			UserName:     "John",
			PlaylistID:   2,
			PlaylistName: "POP 2022",
		},
		{
			UserID:       1,
			UserName:     "John",
			PlaylistID:   3,
			PlaylistName: "Rock",
		},
	}

	playlistTracks := []model.PlaylistTrack{
		{
			PlaylistID:   3,
			PlaylistName: "Rock",
			TrackID:      6,
			TrackName:    "Music F",
			TrackArtist:  "Jane",
		},
	}

	Describe("Fetch user playlists", func() {
		It("should return user playlists", func() {
			db, err := sql.Open("sqlite3", "./music-app.db")
			Expect(err).To(BeNil())

			playListRepo := NewPlaylistRepository(db)
			playlists, err := playListRepo.FetchUserPlaylists(1)
			Expect(err).To(BeNil())
			Expect(playlists).To(Equal(usersPlaylist))
		})
	})

	Describe("Create user playlist", func() {
		It("should create user playlist", func() {
			db, err := sql.Open("sqlite3", "./music-app.db")
			Expect(err).To(BeNil())

			playListRepo := NewPlaylistRepository(db)
			playlist := model.Playlist{
				Name:      "Lo-fi",
				UserID:    1,
				CreatedAt: "2020-01-01 00:00:00",
			}
			err = playListRepo.CreatePlaylist(playlist)
			Expect(err).To(BeNil())

			playlists, err := playListRepo.FetchUserPlaylists(1)
			Expect(err).To(BeNil())
			Expect(len(playlists)).To(Equal(4))
		})
	})

	Describe("Update user playlist", func() {
		It("should update user playlist", func() {
			db, err := sql.Open("sqlite3", "./music-app.db")
			Expect(err).To(BeNil())

			playListRepo := NewPlaylistRepository(db)
			playlist := model.Playlist{
				ID:        1,
				Name:      "Chill Hits Music",
				UserID:    1,
				CreatedAt: "2020-01-01 00:00:00",
			}
			err = playListRepo.UpdateUserPlaylist(int64(1), playlist)
			Expect(err).To(BeNil())

			playlists, err := playListRepo.FetchUserPlaylists(1)
			Expect(err).To(BeNil())

			Expect(playlists[0].PlaylistName).To(Equal("Chill Hits Music"))
		})
	})

	Describe("Fetch playlist track", func() {
		It("should return playlist track", func() {
			db, err := sql.Open("sqlite3", "./music-app.db")
			Expect(err).To(BeNil())

			playListRepo := NewPlaylistRepository(db)

			playlistTracksRes, err := playListRepo.FetchPlaylistTrack(3)
			Expect(err).To(BeNil())
			Expect(playlistTracksRes).To(Equal(playlistTracks))
		})
	})
})
