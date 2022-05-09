package main

import "fmt"

// Nah setelah itu ada sistem baru nih yang masuk namanya Turkey

// sebuah Turkey memiliki interface seperti ini

type Turkey interface {
	Gobble()
	Fly()
}

// Lalu ada WildTurkey untuk implementasi concrete nya
type WildTurkey struct {
}

func (w WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

// Nah karena sebuah turkey hanya bisa Fly() dalam jarak yang pendek
func (w WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}
