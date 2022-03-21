package main

import "fmt"

func main() {
	// langsung kasih nilai ke dua variabel
	sumResult, multiplyResult := calculate(4, 4)
	fmt.Printf("%d dan %d\n", sumResult, multiplyResult)
	fmt.Println(calculate(5, 5))
}

func calculate(angka1, angka2 int) (int, int) {
	sumResult := angka1 + angka2
	multiplyResult := angka1 * angka2
	return sumResult, multiplyResult
}
