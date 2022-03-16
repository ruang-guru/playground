package main

import (
	"fmt"
)

// Check Point:
// Repetition with Condition
// Cases Study: Pada sebuah gudang di toko beras terdapat stok beras yang terbatas dengan kualitas beras yang berbeda-beda.
// 				Jadi pembeli hanya dibatasi untuk bisa membeli hanya 1 kilo per kualitas.
// 				- Kualitas SUPER hanya dapat dibeli oleh 5 orang pertama.
//              - Kualitas MEDIUM bisa dibeli oleh 5 Orang Berikutnya.
//              - Sisanya hanya bisa membeli beras dengan kualitas LOW
// - Input: Masukkan jumlah antrian : jumlah antrian
// - Output: Deretan nilai dengan tanda: Antrian [urutan] membeli 1kg beras dengan kualitas [SUPER] | [MEDIUM] | [LOW]

// Contoh
// - Input: Masukkan jumlah antrian : 15
// - Output:
// Antrian 1 membeli 1kg beras dengan kualitas [SUPER]
// Antrian 2 membeli 1kg beras dengan kualitas [SUPER]
// Antrian 3 membeli 1kg beras dengan kualitas [SUPER]
// Antrian 4 membeli 1kg beras dengan kualitas [SUPER]
// Antrian 5 membeli 1kg beras dengan kualitas [SUPER]
// Antrian 6 membeli 1kg beras dengan kualitas [MEDIUM]
// Antrian 7 membeli 1kg beras dengan kualitas [MEDIUM]
// Antrian 8 membeli 1kg beras dengan kualitas [MEDIUM]
// Antrian 9 membeli 1kg beras dengan kualitas [MEDIUM]
// Antrian 10 membeli 1kg beras dengan kualitas [MEDIUM]
// Antrian 11 membeli 1kg beras dengan kualitas [LOW]
// Antrian 12 membeli 1kg beras dengan kualitas [LOW]
// Antrian 13 membeli 1kg beras dengan kualitas [LOW]
// Antrian 14 membeli 1kg beras dengan kualitas [LOW]
// Antrian 15 membeli 1kg beras dengan kualitas [LOW]

func main() {
	var sizeQueue int
	fmt.Printf("Masukkan jumlah antrian : ")
	fmt.Scan(&sizeQueue)

	// TODO: answer here
}
