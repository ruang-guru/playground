package main

import (
	"bytes"
	"html/template"
	"log"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
	// TODO: answer here
	total := 0
	for _, val := range leaderboard.Users {
		total += val.Score
	}
	return total
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here

	funcMap := template.FuncMap{
		"sum": CalculateScore,
	}

	data := map[string]interface{}{
		"leaderboard": leaderboard,
	}

	textTemplate = "{{range .leaderboard.Users}}{{.Name}}: {{.Score}}{{end}}Total Score: {{sum .leaderboard}}"

	tmp, err := template.New("template_1").Funcs(funcMap).Parse(textTemplate)

	if err != nil {
		log.Fatal(err)
	}

	b := new(bytes.Buffer)

	tmp.Execute(b, data)

	return b.Bytes(), nil
}
