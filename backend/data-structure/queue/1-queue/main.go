// Ada n orang dalam antrian untuk membeli tiket, dimana orang ke-0 berada di barisan paling depan dan orang (n - 1) berada di barisan paling belakang.
// Diberikan array integer dengan nama "tickets" dengan panjang n di mana jumlah tiket yang ingin dibeli orang ke-i adalah tiket[i].
// Setiap orang membutuhkan waktu tepat 1 detik untuk membeli tiket.
// Seseorang hanya dapat membeli 1 tiket pada satu waktu dan harus kembali ke akhir baris jika ingin membeli lebih dari 1 tiket.
// Jika seseorang sudah tidak ingin membeli tiket lagi, orang tersebut akan meninggalkan barisan.
// Kembalikan waktu yang dibutuhkan orang di posisi k untuk menyelesaikan pembelian tiket.
//
// Contoh 1
// Input: tickets = [2,3,2], k = 2
// Output: 6
// Explanation:
// - Pada sesi pertama, semua orang dalam antrian membeli tiket dan antriannya menjadi [1, 2, 1].
// - Pada sesi kedua, semua orang dalam antrian membeli tiket dan antriannya menjadi [0, 1, 0].
// - Orang di posisi (indeks) 2 telah berhasil membeli 2 tiket dan butuh waktu 3+3 = 6 detik.
//
// Contoh 2
// Input: tickets = [5,1,1,1], k = 0
// Output: 8
// Explanation:
// - Pada sesi pertama, semua orang dalam antrean membeli tiket dan antreannya menjadi [4, 0, 0, 0].
// - Dalam 4 sesi berikutnya, hanya orang di posisi (indeks) 0 yang membeli tiket.
// - Orang yang berada di posisi (indeks) 0 telah berhasil membeli 5 tiket dan membutuhkan waktu 4+1+1+1+1= 8 detik.

package main

import (
	"fmt"
	"math"
)

func main() {
	var tickets = []int{2, 3, 3}
	var k = 2
	fmt.Println(timeRequiredToBuy(tickets, k))
}

func timeRequiredToBuy(tickets []int, k int) int {
	val := float64(tickets[k])
	res := 0

	for i := 0; i < len(tickets); i++ {
		res += int(math.Min(float64(tickets[i]), val))
	}
	return res
}
