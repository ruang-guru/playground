package main

import "fmt"

func main() {
	//fungsi counter akan menerima (x int) dan mengembalikan fungsi
	//fungsi yang dikembalikan akan melakukan decrement nilai parameter x ketika dipanggil dan
	//mengembalikan nilai parameter x

	counter := func(x int) func() int {
<<<<<<< HEAD
		// TODO: answer here
=======
		//beginanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
		return func() int {
			x--
			return x
		}
<<<<<<< HEAD
=======
		//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	}
	decrement := counter(5)
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}
