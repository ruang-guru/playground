package main

import "fmt"

func main() {
	// Pada bahasa Go tidak ada keyword while, untuk melakuka looping while, kita harus menggunakan looping for
	// disini kita akan mengenal keyword `continue` dan `break`
	// break digunakan untuk menghentikan looping jika suatu kondisi terpenuhi
	// karena jika while(for) tidak mempunyai exit plan maka dia akan berjalan terus
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	// contoh for loop dengan continue
	// program dibawah ini akan melakukan skip jika kondisi terpenuhi
	// kalian juga bisa menaruk exit plan (break) seperti ini
	k := 0
	for k <= 10 {
		k++
		if k%2 == 0 {
			continue
		}
		fmt.Print(k)
	}
}
