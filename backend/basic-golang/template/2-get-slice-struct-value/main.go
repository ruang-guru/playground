package main

import (
	"fmt"
	"os"
	"text/template"
)

//Membuat struct Student
//Memiliki field Name dan Age
type Student struct {
	Name string
	Age  int
}

// main function
func main() {
	stds := []Student{{Name: "Rogu", Age: 16}, {Name: "Dafa", Age: 10}, {Name: "Lulu", Age: 10}}

	// "New" membuat template baru
	// dengan nama Template_2
	tmp1 := template.New("Template_2")

	// "Parse" parses string dan set action untuk mengambil nilai dari Field Name dan Age
	// tambahkan action range untuk get value dari struct slice
	tmp1, _ = tmp1.Parse("{{range .}}Hello {{.Name}}, usia kamu {{.Age}}!{{end}}\n")

	// standard output untuk menampilkan
	// string beserta value dari struct
	err := tmp1.Execute(os.Stdout, stds)

	// print error jika ada
	if err != nil {
		fmt.Println(err)
	}
}
