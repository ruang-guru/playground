package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Kita juga bisa menggunakan `if` sintax didalam template.
// Misal kita ingin melakukan pengecekan apakah `FieldName` yang akan diambil mempunyai value atau hanya empty string.
// Jika `FirstName` mempunyai value maka output nya adalah `Halo Rogu!`. Lalu jika `FirstName` tidak ada valuenya maka tidak akan ada output apapun.
// Lalu bagaimana untuk melakukan operational `if else` condition? kita bisa lakukan seperti berikut:

func main() {
	textTemplate := `{{if (gt .Age 16) }} 
Halo {{ .FirstName }}, Selamat usia kamu sudah memenuhi syarat untuk membuat SIM!
{{ else }} 
Halo {{ .FirstName }}, Mohon maaf usia kamu masih belum memenuhi syarat.
{{ end }}`

	type User struct {
		FirstName string
		Age       int
	}
	u := User{"Rogu", 17} // set value dari struct user dan disimpan ke variable u
	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, u)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b.Bytes()))
}
