//jangan ditunjukkan ke peserta
//mungkin dikerjakan jika waktu cukup aja
package main

import "fmt"

//fungsi calculate akan mengembalikan hasil melakukan perhitungan berikut
//penjumlahan, pengurangan,perkalian,dan pembagian
func main() {
	sumResult, subtractResult, multiplyResult, divideResult := calculate(4, 4)
	fmt.Println("hasil penjumlahan", sumResult)
	fmt.Println("hasil pengurangan", subtractResult)
	fmt.Println("hasil perkalian", multiplyResult)
	fmt.Println("hasil pembagian", divideResult)
	fmt.Println(calculate(5, 5))
}

func calculate(Number1, Number2 int) (sumResult, subtractResult, multiplyResult, divideResult int) {
<<<<<<< HEAD
	// TODO: answer here
	fmt.Printf("\nAngka yang dimasukan = %d dan %d\n\n", Number1, Number2)
=======
	//beginanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	sumResult = Number1 + Number2
	subtractResult = Number1 - Number2
	multiplyResult = Number1 * Number2
	divideResult = Number1 / Number2
<<<<<<< HEAD
=======
	//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	return
}
