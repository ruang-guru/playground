package main

import "fmt"

// "If its quack like a duck its a duck"

// let say kita mendefine interface seperti ini

type Duck interface {
	Quack()
	Fly()
}

// Nah disini kita membuat concrete dari Duck tersebut katakan MallardDuck

type MallardDuck struct {
}

func (m MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (m MallardDuck) Fly() {
	fmt.Println("I'm flying")
}
