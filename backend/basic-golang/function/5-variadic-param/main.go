package main

import "fmt"

func main() {
	createSentence("halo")
	createSentence("halo", "selamat siang")
	createSentence("halo", "andi", "dan", "bagus")
	createSentence("mencoba", "variadic", "param", "pada", "go")
}

func createSentence(words ...string) {
	result := ""
	for _, word := range words {
		result += " " + word
	}
	fmt.Println("hasil membuat kalimat", result)
}
