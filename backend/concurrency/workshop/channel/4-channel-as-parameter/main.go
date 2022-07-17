package main

import "fmt"

func main() {
	c := make(chan int)
	a := 4
	b := 5

	go multiply(a, b, c)

	fmt.Printf("hasil perkalian %d dan %d: %d\n", a, b, <-c) // menerima data dari channel c
}

//channel sebagai parameter
func multiply(a, b int, c chan int) {
	fmt.Printf("kirim hasil perkalian ke channel\n")
	c <- a * b //hasil perkalian dikirim ke channel c
}
