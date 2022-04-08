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

func (c cat) String() string {
	return fmt.Sprintf("%v my name is %v", c.sound, c.name)
}
