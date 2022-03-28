package main

import (
	"fmt"
	"strings"
)

// Golang memiliki banyak standard library yang dapat digunakan. Standard library ini merupakan sekumpulan package yang memiliki fungsi dan tujuan yang berbeda-beda.
// Contoh package yang sering kita gunakan adalah `fmt` yang dapat digunakan untuk melakukan formatting string dan mencetak variabel ke dalam output yang diinginkan.
// Untuk list dari standard library kamu bisa akses di https://pkg.go.dev/strings@go1.18

func main() {
	word := "Hello World"
	fmt.Println(strings.Contains(word, "Hello"))
}
