package main

import "fmt"

// Lanjutan nomor 2
// Sehabis menambahkan "Olleh" pada slice tersebut coba ubah nilai "World" menjadi "Marcus"
// dan "Olleh" menjadi "Aurelius"
func main() {
	slice := []string{"Hello", "World"}

	slice = append(slice, "Olleh")

	slice[1] = "Marcus"
	slice[len(slice)-1] = "Aurelius"

	for index, value := range slice {
		fmt.Println(index, value)
	}
}
