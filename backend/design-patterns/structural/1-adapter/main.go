package main

import "fmt"

// run this with `go run .`
func main() {
	duck := MallardDuck{}
	turkey := WildTurkey{}

	turkeyAdapter := &TurkeyAdapter{
		Turkey: turkey,
	}

	fmt.Println("The Turkey says...")
	turkey.Gobble()
	turkey.Fly()

	fmt.Println("\nThe Duck says...")
	testDuck(duck)

	// Disini kita bisa lihat bahwa kita sudah menggunakan adapter
	// Untuk turkey agar bisa menggunakan function testduck yang menerima interface Duck
	fmt.Println("\nThe TurkeyAdapter says...")
	testDuck(turkeyAdapter)
}

func testDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}
