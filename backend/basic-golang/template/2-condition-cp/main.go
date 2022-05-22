package main

import (
	"bytes"
	"html/template"
	"log"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan condition pada template.
// Lengkapi function ExecuteToByteBuffer dan variabel textTemplate sehingga memiliki keluaran:
// 1. Jika saldo sama dengan 0, cetak "Akun Tony dengan Nomor 1002321 tidak memiliki saldo."
// 2. Jika saldo lebih dari 0, cetak "Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	textTemplate = "Akun {{.Name}} dengan Nomor {{.Number}} {{if gt .Balance 0}}memiliki saldo sebesar ${{.Balance}}{{else}}tidak memiliki saldo{{end}}."

	b := new(bytes.Buffer)

	template, err := template.New("template_1").Parse(textTemplate)
	if err != nil {
		log.Fatal(err)
	}

	err = template.Execute(b, account)
	if err != nil {
		log.Fatal(err)
	}

	return b.Bytes(), nil
}
