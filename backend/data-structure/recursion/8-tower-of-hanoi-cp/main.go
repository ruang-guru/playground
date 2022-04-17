// ===============>>>> Game Tower of Hanoi: <<<<====================
// Tower of Hanoi adalah teka-teki matematika di mana kita memiliki tiga batang dan n piringan.
// Tujuan dari teka-teki ini adalah untuk memindahkan seluruh tumpukan ke batang lain, dengan mematuhi aturan sederhana berikut:

// 1. Hanya satu disk yang dapat dipindahkan pada satu waktu.
// 2. Setiap langkah terdiri dari mengambil disk atas dari salah satu tumpukan dan menempatkannya di atas tumpukan lain yaitu disk hanya dapat dipindahkan jika disk paling atas pada tumpukan.
// 3. Tidak ada disk yang dapat ditempatkan di atas disk yang lebih kecil.

// Contoh:
// Misalkan batang 1 = 'A', batang 2 = 'B', batang 3 = 'C'.
// Contoh dengan 2 disk :
// Langkah 1 : Geser disk pertama dari 'A' ke 'B'.
// Langkah 2: Geser disk kedua dari 'A' ke 'C'.
// Langkah 3: Pindahkan disk pertama dari 'B' ke 'C'.

// Contoh dengan 3 disk :
// Langkah 1 : Geser disk pertama dari 'A' ke 'C'.
// Langkah 2: Geser disk kedua dari 'A' ke 'B'.
// Langkah 3: Pindahkan disk pertama dari 'C' ke 'B'.
// Langkah 4: Geser disk ketiga dari 'A' ke 'C'.
// Langkah 5: Geser disk pertama dari 'B' ke 'A'.
// Langkah 6: Geser disk kedua dari 'B' ke 'C'.
// Langkah 7: Geser disk pertama dari 'A' ke 'C'. (Perhatikan celahnya)

// Polanya adalah :
//  - Geser 'n-1' disk dari 'A' ke 'B', menggunakan C.
//  - Geser disk terakhir dari 'A' ke 'C'.
//  - Geser disk 'n-1' dari 'B' ke 'C', menggunakan A.
// Lihat gambar berikut: https://media.geeksforgeeks.org/wp-content/uploads/tower-of-hanoi.png

// Input : 2
// Output : Disk 1 dipindahkan dari A ke B
//          Disk 2 dipindahkan dari A ke C
//          Disk 1 dipindahkan dari B ke C

// Input : 3
// Output : Disk 1 dipindahkan dari A ke C
//          Disk 2 dipindahkan dari A ke B
//          Disk 1 dipindahkan dari C ke B
//          Disk 3 dipindahkan dari A ke C
//          Disk 1 dipindahkan dari B ke A
//          Disk 2 dipindahkan dari B ke C
//          Disk 1 dipindahkan dari A ke C

// Cobalah game Tower of Hanoi di halaman ini: https://www.mathsisfun.com/games/towerofhanoi.html
// Tonton Video penjelasan dari Tower of Hanoi: https://www.youtube.com/watch?v=YstLjLCGmgg

// Referensi:
// https://en.wikipedia.org/wiki/Tower_of_Hanoi

// Buatlah solusi dari Tower of Hanoi dengan menggunakan rekursif.:
// Replace pada TODO dengan kode yang sesuai.
package main

import "fmt"

var Solution []string

func TowerOfHanoi(n int, fromRod string, auxRod string, toRod string) {
	if n == 0 {
		return
	}

	// TODO: answer here
	Solution = append(Solution, fmt.Sprintf("Move disk %d from rod %s to rod %s", n, fromRod, toRod))
	// TODO: answer here
}

func main() {
	n := 4 // Number of disks
	TowerOfHanoi(n, "A", "B", "C")
	for i := 0; i < len(Solution); i++ {
		fmt.Println(Solution[i])
	}

}
