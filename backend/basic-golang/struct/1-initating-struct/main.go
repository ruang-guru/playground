package main

import "fmt"

//membuat struct Person
//memiliki field/atribut Name dan Age
type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	p.Name = "widi"
	p.Age = 42
	fmt.Println(p)

	p2 := Person{Name: "woody", Age: 10}
	fmt.Println(p2)

	//akses field
	fmt.Println("halo nama saya", p2.Name)
}
