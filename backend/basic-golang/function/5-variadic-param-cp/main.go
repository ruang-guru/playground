package main

import "fmt"

//fungsi printWord akan melakukan print satu persatu nilai parameter yang diterimanya
//contoh nilai parameter yang diterima
//("selamat","pagi","siang",sore)
//outputnya
//selamat
//pagi
//siang
//sore
func main() {
	printWord("halo")
	printWord("halo", "selamat siang")
	printWord("halo", "andi", "dan", "bagus")
	printWord("mencoba", "variadic", "param", "pada", "go")
}

<<<<<<< HEAD
// TODO: answer here
func printWord(key ...string){

	for _, item := range key {
		fmt.Println(item)
	}
}
=======
//beginanswer
func printWord(words ...string) {
	for _, word := range words {
		fmt.Println(word)
	}
}

//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
