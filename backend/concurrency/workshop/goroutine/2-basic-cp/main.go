package main

import (
	"fmt"
	// TODO: answer here
)

func multiply(a, b int, called *bool) {
	fmt.Printf("hasil perkalian %d dan %d: %d\n", a, b, a*b)
	*called = true
}

//jalankan fungsi multiply dengan goroutine
//hint: gunakan sesuatu yang blocking
func start(multiplyCalled *bool) {
	go multiply(4, 5, multiplyCalled)
	// TODO: answer here
}
