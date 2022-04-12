package main

import "fmt"

type cat struct {
	name  string
	sound string
}

func main() {
	cat1 := cat{name: "dorito", sound: "meow"}
	fmt.Println(cat1)
}
