package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Student struct {
	Name string
	Age  int
	Id   string
}

// main function
func main() {
	// defining struct instance
	std1 := Student{Name: "Rogu", Age: 14, Id: "128890"}

	// Parsing file html
	// file ini berada di satu folder dengan main
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("ParseFiles error: %s", err)
	}

	// standard output untuk print string dan struct value
	if err := t.Execute(os.Stdout, std1); err != nil {
		// print error jika ada
		fmt.Println(err)
	}
}
