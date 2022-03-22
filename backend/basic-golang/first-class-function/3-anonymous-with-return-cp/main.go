package main

import "fmt"

func main() {
	// mengembalikan string selamat sore dengan anonymous function
	goodAfternoon := func() string {
<<<<<<< HEAD
		// TODO: answer here
		return "Selamt sore"
=======
		//beginanswer
		return "Selamat sore"
		//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
	}()

	fmt.Println(goodAfternoon)
}
