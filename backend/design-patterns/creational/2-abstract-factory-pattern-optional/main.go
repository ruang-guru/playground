package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/design-patterns/creational/2-abstract-factory-pattern-optional/abstract-factory/samsung"
)

func main() {
	samsung := samsung.Samsung{}
	smartphone := samsung.CreateSmartphone()
	fmt.Println(smartphone.Price())
}
