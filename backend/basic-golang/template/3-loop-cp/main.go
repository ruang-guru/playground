package main

import (
	"bytes"
	"html/template"
	"log"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untubvk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
	Name  string
	Email string
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here

	textTemplate = "{{range .Users}}Peringkat ke-{{.Rank}}: {{.Name}}{{end}}"

	tmp, err := template.New("template_1").Parse(textTemplate)

	if err != nil {
		log.Fatal(err)
	}

	b := new(bytes.Buffer)

	err = tmp.Execute(b, leaderboard)

	return b.Bytes(), err
}
