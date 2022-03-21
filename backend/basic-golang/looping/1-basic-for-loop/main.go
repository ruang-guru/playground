package main

import "fmt"

// sebenarnya kalian sudah mendapatkan eksposur looping dari materi materi sebelumya
// tetapi disinin akan kita bahas lebih lanjut tentang looping

func main() {
	// contoh looping for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// contoh looping for dengan kondisi
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// contoh looping while
	// Pada bahasa Go tidak ada keyword while, untuk melakuka looping while, kita harus menggunakan looping for
	// contoh looping while
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
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}
}
