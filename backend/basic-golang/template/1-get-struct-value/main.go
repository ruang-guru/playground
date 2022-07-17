package main

import (
	"bytes"
	"fmt"
	"log"
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
	buff := new(bytes.Buffer)
	std := Student{Name: "Rogu", Age: 16}

	// "New" membuat template baru
	// dengan nama Template_1
	tmp1 := template.New("Template_1")

	// "Parse" parses string dan set action untuk mengambil nilai dari Field Name dan Age
	tmp1, err := tmp1.Parse("Hello {{.Name}}, usia kamu {{.Age}}!")
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}

	// write string template dan value dari struct kedalam buffer
	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}

	// print buffer dalam format string
	fmt.Println(buff.String())
}
