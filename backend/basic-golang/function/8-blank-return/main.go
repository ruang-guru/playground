package main

import "fmt"

func main() {
	sumResult, multiplyResult := calculate(4, 4)
	fmt.Printf("%d dan %d\n", sumResult, multiplyResult)
	fmt.Println(calculate(5, 5))
}

// karena kita sudah menentukan variabel apa yang akan dikembalikan
// kita bisa langsung memanggil return
func calculate(number1, number2 int) (sumResult, multiplyResult int) {
	sumResult = number1 + number2
	multiplyResult = number1 * number2
	return
}
