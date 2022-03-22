//jangan ditunjukkan ke peserta
//mungkin dikerjakan jika waktu cukup aja
package main

import "fmt"

// fungsi ini digunakan untuk menambahkan point
// fungsi ini merupakan closure
func points(base int) func(x int) int {
<<<<<<< HEAD
	// TODO: answer here
	value := base
	return func(x int) int {
		value += x
		return value
=======
	//beginanswer
	return func(points int) int {
		base += points
		return base
		//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	}
}

func main() {
	toni := points(100) // base value 100
	tono := points(100)
	fmt.Println(toni(2))   // jadi 102
	fmt.Println(toni(3))   // 105
	fmt.Println(toni(100)) // 205
	fmt.Println(tono(2))   // 102
}
