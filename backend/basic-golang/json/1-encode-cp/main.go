package main

import "encoding/json"

// Dari contoh yang telah diberikan, cobalah untuk melakukan encode struct menjadi json.
// Lengkapi function EncodeToJson agar dapat mengembalikan nilai byte hasil dari encode objek Leaderboard.
// Modifikasi struct UserRank sehingga field Name menjadi name ,field Rank menjadi rank, dan field Email tidak ikut untuk diencode.

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
}

type Leaderboard struct {
	Users []*UserRank
}

func EncodeToJson(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
}
