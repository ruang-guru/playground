package main

import "fmt"

// Lanjutan nomor 2
// Sehabis menambahkan "Olleh" pada slice tersebut coba ubah nilai "World" menjadi "Marcus"
// dan "Olleh" menjadi "Aurelius"
func main() {
	slice := []string{"Hello", "World"}

	// Dibawah ini adalah jawaban nomor 2: silahkan kalian copy paste dari jawaban nomor 2
	slice = append(slice, "Olleh")

	// Dibawah ini adalah jawaban nomor 3 silahkan kalian isi
	slice[1] = "Marcus"
	slice[2] = "Aurelius"

	for _, value := range slice {
		fmt.Println(value)
	}
}
