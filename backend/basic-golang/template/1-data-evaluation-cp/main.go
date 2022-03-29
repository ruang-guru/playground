package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, cobalah untuk mempraktikkan data evaluation pada template.
// Lengkapi function ExecuteToByteBuffer sehingga objek dari struct Account dapat ter-render ke dalam template.
// Gunakan bytes.Buffer sebagai io.Writer pada template.
// Lengkapi juga textTemplate, sehingga variabel dari objek Account dapat ter-render ke dalam template.
// Contoh:
// acc := {Name: "Tony", Number: "1002321", Balance: 1000}
// Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000.

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
}
