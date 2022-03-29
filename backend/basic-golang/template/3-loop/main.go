package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Kita juga bisa melakukan iterasi didalam template menggunakan syntax `range` action.

func main() {
	textTemplate := `{{range . }}
First Name: {{ .FirstName }}
Age: {{ .Age }}
{{else}} 
Invalid "struct" Users harus berupa array!
{{end}}`

	type User struct {
		FirstName string
		Age       int
	}
	users := []User{
		{
			FirstName: "Tony",
			Age:       40,
		},
		{
			FirstName: "Clint",
			Age:       35,
		},
		{
			FirstName: "Natasha",
			Age:       34,
		},
	}
	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, users)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b.Bytes()))
}
