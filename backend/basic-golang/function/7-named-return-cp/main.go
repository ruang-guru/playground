package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan named return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

<<<<<<< HEAD
// TODO: answer here
func square(num1, num2 int) (res1, res2 int) {
	res1 = num1 * num1
	res2 = num2 * num2
	return res1, res2
}
=======
//beginanswer
func square(angka1, angka2 int) (result1 int, result2 int) {
	result1 = angka1 * angka1
	result2 = angka2 * angka2
	return result1, result2
}

//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
