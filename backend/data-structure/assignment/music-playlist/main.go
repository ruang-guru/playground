package main

import "fmt"

// Implementasikan MusicPlayer yang merupakan representasi dari pemutar musik.
// MusicPlayer memiliki atribut Playlist. Playlist merupakan daftar lagu yang tersimpan.
// Sebagai gambaran, seperti pada Spotify, playlist merepresentasikan lagu-lagu yang disimpan dalam satu playlist yang sama.
// Lagu direpresentasikan dalam objek Song yang memiliki dua atribut yaitu Singer dan Title.
// Di Playlist juga terdapat atribut IsRepeatable yang merepresentasikan apakah playlist tersebut akan diputar ulang jika lagu telah habis.

// Lengkapilah tiga method yang ada pada MusicPlayer yaitu AddSong, Play, dan Repeat.
// AddSong merupakan method yang menambahkan lagu pada playlist.
// Play merupakan method yang memutar lagu pada playlist. Pada method ini, cetaklah sebuah wording dengan format `Now playing [Singer] - [Title]`
// 		Sebagai contoh, jika lagu One Direction dengan judul Night Changes diputar, maka akan mengeluarkan output `Now playing One Direction - Night Changes`
// Repeat merupakan method yang memungkinkan playlist diputar ulang jika lagu telah habis.
// Jika playlist dilakukan repeat, maka ketika lagu telah diputar semua maka akan mengulang pada urutan semula.
type Song struct {
	Singer string
	Title  string
}

type Playlist struct {
	Songs        []Song
	IsRepeatable bool
}

func NewPlaylist() Playlist {
	return Playlist{
		Songs:        []Song{},
		IsRepeatable: false,
	}
}

func (pl Playlist) IsEmpty() bool {
	return len(pl.Songs) == 0
}

type MusicPlayer struct {
	Playlist Playlist
}

func NewMusicPlayer() MusicPlayer {
	return MusicPlayer{
		Playlist: NewPlaylist(),
	}
}

func (mp *MusicPlayer) AddSong(song Song) {
	// TODO: answer here
}

func (mp *MusicPlayer) Play() string {
	// TODO: answer here
}

func (p *Playlist) Repeat() {
	// TODO: answer here
}
