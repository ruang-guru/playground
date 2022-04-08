// Queue Insertion and Deletion
// Queue adalah data structure yang menyimpan elemen-elemen yang berurutan.
// Queue ini dapat menyimpan elemen-elemen yang berurutan, yaitu elemen yang ditambahkan pada bagian depan dan elemen yang ditambahkan pada bagian belakang.

// Dalam kasus ini, kita akan membuat queue dengan tipe data string.
// Terdapat tiga fungsi utama yang akan digunakan untuk mengoperasikan queue:
// insert, delete, dan display.
// 1. Fungsi insert akan menambahkan elemen baru pada queue berurutan dari kiri ke kanan.
//    - Contoh: [Dion nil nil nil nil] -> [Dion Gina nil nil nil] -> [Dion Gina Juno nil nil] -> [Dion Gina Juno Bagas nil]
// 2. Fungsi delete akan menghapus elemen pada queue berurutan dari kiri ke kanan.
//    - Contoh: [nil Gina Juno Bagas nil] -> [nil nil Juno Bagas nil] -> [nil nil nil Bagas nil] -> [nil nil nil nil nil]
// 3. Fungsi display akan menampilkan elemen pada queue.

package main

import (
	"fmt"
)

// r - rear
// n - length
// f - front
// q - queue

type queue []string

var r, n, f int
var q queue

// Fungsi ini akan dipanggil pertama kali saat program dijalankan.
// Fungsi ini akan menginisialisasi queue dengan nilai awal yaitu nil dengan maximal panjang queue 5.
func init() {
	r = -1
	f = -1
	n = 4
	q = make(queue, 5)
	for i := range q {
		q[i] = "nil"
	}
}

func main() {
	i := 0
	var choice int
	for i == 0 {
		fmt.Printf("\x1bc") // clear screen
		display()
		fmt.Println("Main Menu Queue:")
		fmt.Println("1. INSERT")
		fmt.Println("2. DELETE")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			q.insert()
			fmt.Println("\nValues after insertion FRONT:", f, "REAR:", r)

			pause()
		case 2:
			q.delete()
			fmt.Println("\nValues after deletion FRONT:", f, "REAR:", r)

			pause()
		case 3:
			fmt.Println("\nFRONT:", f, "REAR:", r)

			pause()
		case 4:
			i = 1
		default:
			fmt.Println("\nCommand not recognized")
		}
	}
}

func (q queue) insert() {
	// check for overflow
	if r >= n {
		fmt.Println("\n-- Overflow --")
		return
	}
	// increment rear pointer
	r++
	// insert the element
	var y string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s\n", &y)
	q[r] = y

	// set front pointer
	if f == -1 {
		f = 0
	}
}

func (q queue) delete() {
	// check for underflow
	if f == -1 {
		fmt.Println("\n-- Underflow --")
		return
	}
	// delete element
	y := q[f]
	q[f] = "nil"

	// check if queue is empty
	if f == r {
		f = -1
		r = -1
	} else {
		f++
	}
	// print deleted element
	fmt.Println("Element deleted", y)
}

func display() {
	for i := range q {
		fmt.Print(q[i], " ")
	}
	fmt.Println("")
}

func pause() {
	var input string
	fmt.Print("Press enter to continue...")
	fmt.Scanf("%s\n", &input)
}
