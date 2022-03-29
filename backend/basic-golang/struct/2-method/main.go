package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//Method Eat akan print apa yang Person makan
func (p Person) Eat(food string) {
	fmt.Printf("%s makan %s\n", p.Name, food)
}

//Method Sleep akan print Person tidur
func (p Person) Sleep() {
	fmt.Printf("%s tidur\n", p.Name)
}

func main() {
	var p Person
	p.Name = "putra"
	p.Age = 20
	fmt.Println(p)
	p.Eat("nasi goreng")
	p.Sleep()
}
