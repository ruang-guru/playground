package main

import "fmt"

// Demo:
// Sum Average
// Input
//    - Jumlah Anggka: angka
//    - Angka ke-1: Angka ke-2: angka angka ...
// Output
// 	  - Angka ke-3: Angka ke-[i]: Angka ke-[i] ... Jumlah: [Jumlah Angka]
//    - Rata-rata: nilai rata-tata

// Contoh
// Input:
//    - Jumlah Angka: 10
//    - Angka ke-1: Angka ke-2: 2 3 4 5 6 7 8 9 5 8
// Output:
//    - Angka ke-3: Angka ke-4: Angka ke-5: Angka ke-6: Angka ke-7: Angka ke-8: Angka ke-9: Angka ke-10: Jumlah: 49
//    - Rata-rata: 4.9

func main() {
	var n int
	fmt.Print("Jumlah Angka: ")
	fmt.Scanf("%d", &n)

	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Angka ke-", i+1, ": ")
		var num int
		fmt.Scanf("%d", &num)
		numbers[i] = num
	}

	var sum int
	for _, num := range numbers {
		sum += num
	}

	var average float64
	average = float64(sum) / float64(n)

	fmt.Println("Jumlah:", sum)
	fmt.Println("Rata-rata:", average)
}
