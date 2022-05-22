package main

import "encoding/json"

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
	Name  string
	Email string `json:"-"`
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here
	leaderboard := Leaderboard{}

	err := json.Unmarshal(jsonData, &leaderboard)
	return leaderboard, err
}
