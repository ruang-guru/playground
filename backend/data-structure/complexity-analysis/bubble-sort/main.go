// Bubble Sort

// Bubble sort didasarkan pada gagasan untuk membandingkan pasangan elemen yang berdekatan secara berulang dan kemudian menukar posisi mereka jika ada dalam urutan yang salah.
// Asumsikan bahwa A[] dalah array yang tidak disortir dari n elemen. Array ini perlu diurutkan dalam urutan menaik.

// Contoh:
// Mari kita coba memahami kode semua dengan sebuah contoh: A[] = { 7, 4, 5, 2}

// Image:
// Untuk memahami solusi sort diatas dengan Bubble Sort, bisa diperhatikan gambar ini: https://he-s3.s3.amazonaws.com/media/uploads/2682167.png

// Penjelasan:
// - Pada langkah 1. 7 dibandingkan dengan 4. Sejak 7 > 4, dipindahkan ke depan 4. Karena semua elemen lainnya memiliki nilai yang lebih rendah dari pada 7, 7 dipindahkan ke akhir array.
//   Sekarang arraynya adalah A[] = {4, 5, 2, 7}.

// - Pada langkah 2. 4 dibandingkan dengan 5. Sejak 5 > 4 dan keduanya 4 dan 5 berada dalam urutan menaik, elemen-elemen ini tidak ditukar.
//   Namun, ketika 5 dibandingkan dengan 2,5 > 2, dan elemen-elemen ini dalam urutan menurun. Karena itu, 5 dan 2 ditukar.
//   Sekarang arraynya adalah A[] = {4, 2, 5, 7}.

// - Pada langkah 3, elemen 4 dibandingkan dengan 2. Sejak 4 > 2 dan unsur-unsur dalam urutan menurun, 4 dan 2 ditukar.
//   Array yang diurutkan adalah A[] = {2, 4, 5, 7}.

// Visualisasi:
// Contoh Visualisasi Bubble Sort dapat dilihat di sini: https://www.hackerearth.com/practice/algorithms/sorting/bubble-sort/visualize/

// Time Complexity:
// Time Complexity bubble sort adalah O(n2) / O(n*n) dalam kasus terburuk dan rata-rata, karena seluruh array perlu diulang untuk setiap elemen.

// Space Complexity: O(1)

package main

import "fmt"

func main() {
	sample := []int{3, 4, 5, 2, 1}
	bubbleSort(sample)
	sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	bubbleSort(sample)
}

func bubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("\nAfter Bubble Sorting")
	for _, val := range arr {
		fmt.Println(val)
	}
}
