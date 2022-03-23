package main

import "fmt"

// Pointer memiliki zero value berupa nil. Ketika kita mencoba mengakses atribut dari nil, maka program akan terjadi panic.

type Lamp struct {
	State string
}

func TurnOn(lamp *Lamp) {
	// Mengecek apakah lamp merupakan nil
	if lamp == nil {
		fmt.Println("Lamp is nil")
		return
	}
	lamp.State = "On"
}

func main() {
	var lamp *Lamp // nil
	TurnOn(lamp)
}
