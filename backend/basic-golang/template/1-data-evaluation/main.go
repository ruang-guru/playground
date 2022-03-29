package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Data evaluation digunakan untuk menampilkan value dari `struct` yang sudah didefinisikan dalam Golang sintax.
// Untuk menampilkan value dari field yang ada pada struct ini bisa dilakukan dengan cara `{{ .FieldName }}`.

func main() {
	type User struct {
		FirstName string
		Age       int
	}
	u := User{"Rogu", 17} // set value dari struct user dan disimpan ke variable u
	// initiate new template dengan nama "test"
	tmpl, err := template.New("test").Parse("Usia {{.FirstName}} saat ini adalah {{.Age}} tahun.") // ".FirstName" dan ".Age" adalah field name yang ada di struct User
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
