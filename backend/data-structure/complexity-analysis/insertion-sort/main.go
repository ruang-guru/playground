// Insertion Sort didasarkan pada gagasan bahwa satu elemen dari elemen input dikonsumsi dalam setiap iterasi
// untuk menemukan posisinya yang benar, yaitu posisi yang dimilikinya dalam array yang diurutkan.

// Ini mengulangi elemen input dengan menumbuhkan array yang diurutkan pada setiap iterasi.
// Ini membandingkan elemen saat ini dengan nilai terbesar dalam array yang diurutkan.
// Jika elemen saat ini lebih besar, maka ia meninggalkan elemen di tempatnya dan pindah ke elemen berikutnya
// jika tidak menemukan posisi yang benar dalam array yang diurutkan dan memindahkannya ke posisi itu.
// Ini dilakukan dengan menggeser semua elemen, yang lebih besar dari elemen saat ini, dalam array yang diurutkan ke satu posisi di depan

// Contoh:
// Mari kita coba memahami kode semua dengan sebuah contoh: A[] = { 7, 4, 5, 2}

// Image:
// Untuk memahami solusi sort diatas dengan Insertion Sort, bisa diperhatikan gambar ini: https://he-s3.s3.amazonaws.com/media/uploads/46bfac9.png

// Penjelasan:
// Sejaka 7 dalah elemen pertama yang tidak memiliki elemen lain untuk dibandingkan, ia tetap pada posisinya.
// Sekarang ketika bergerak menuju 4,7 adalah elemen terbesar dalam daftar yang diurutkan dan lebih besar dari 4.
// Jadi, pindah 4 ke posisi yang benar yaitu sebelumnya 7. Demikian pula dengan 5, sebagai(elemen terbesar dalam daftar yang diurutkan) lebih besar dari 5,
// kita akan pindahkan 5 ke posisinya yang benar. Akhirnya untuk 2, semua elemen di sisi kiri 2 (daftar yang diurutkan) dipindahkan satu posisi ke depan karena semuanya lebih besar darilaluditempatkan pada posisi pertama.
// Akhirnya, array yang diberikan akan menghasilkan array yang diurutkan.

// Visualisasi:
// Contoh Visualisasi Insertion Sort dapat dilihat di sini: https://www.hackerearth.com/practice/algorithms/sorting/bubble-sort/visualize/

// Time Complexity:
// - Kasus Terburuk - O(n2) / O(n*n) Karena setiap elemen dibandingkan dengan semua elemen lain dalam array yang diurutkan.
// - Kasus Terbaik- O(n) – Ketika array sudah diurutkan

// Space Complexity:
// - Space Complexity dari insertion sort adalah O(1)

package main

import "fmt"

func main() {
	sample := []int{3, 4, 5, 2, 1}
	insertionSort(sample)

	sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	insertionSort(sample)
}

func insertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Println("After Sorting")
	for _, val := range arr {
		fmt.Println(val)
	}
}
