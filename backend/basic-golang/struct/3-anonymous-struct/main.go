package main

import "fmt"

func main() {
	// Membuat objek dengan anonymous struct
	newPerson := struct {
		Name string
		Age  int
	}{
		Name: "putra",
		Age:  20,
	}

	fmt.Printf("orang bernama %s dan berumur %d", newPerson.Name, newPerson.Age)
}

//biasa digunakan untuk satu instance/object saja
