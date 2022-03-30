package main

import (
	"html/template"
	"log"
	"os"
)

// Golang template menyediakan *predefined funtions* seperti `gt` (*greater than*) yang menjalankan fungsi operator `>` seperti pada contoh condition.
// Kita juga bisa buat custom function yang nanti akan di panggil di dalam template.

type SumInput struct {
	A int
	B int
}

// Buat function yang akan di panggil didalam template
func hitungJumlah(a int, b int) int {
	return a + b
}

func main() {
	funcMap := template.FuncMap{
		// key "sum" akan menjadi nama function yang akan dipanggil di template.
		"sum": hitungJumlah,
	}
	// Define simple template untuk test function yang sudah dibuat.
	// sum nama fungsi, .A input untuk parameter pertama dan .B untuk input parameter kedua
	const templateText = `
Result: {{sum .A .B}}
`
	// Buat template, add function map, dan parse template text.
	tmpl, err := template.New("sumTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}
	// Execute template untuk verifikasi output.
	err = tmpl.Execute(os.Stdout, SumInput{10, 20})
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
