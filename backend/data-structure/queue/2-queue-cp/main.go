// Queue Insertion and Deletion
// Queue adalah data structure yang menyimpan elemen-elemen yang berurutan.
// Queue ini dapat menyimpan elemen-elemen yang berurutan, yaitu elemen yang ditambahkan pada bagian depan dan elemen yang ditambahkan pada bagian belakang.

// Dalam kasus ini, kita akan membuat queue dengan tipe data string.
// Terdapat beberapa fungsi yang dapat digunakan untuk mengoperasikan queue.
// 1. Fungsi InsertRear() digunakan untuk menambahkan elemen pada bagian belakang.
//    - Penjelasan: InsertRear("Rear") akan menambahkan elemen "Rear" pada bagian belakang. Jika queue sudah penuh, maka akan mengembalikan error.
//    - Contoh: [nil nil nil nil nil] => [Tino nil nil nil nil nil] => [Tino Rino nil nil nil nil] => ...
// 2. Fungsi InsertFront() digunakan untuk menambahkan elemen pada bagian depan.
//    - Penjelasan: InsertFront("Front") akan menambahkan elemen "Front" pada bagian depan. Jika queue dibagian depan sudah terisi, maka akan mengembalikan error.
//    - Contoh: [nil nil nil nil nil] => [Tino nil nil nil nil].
// 3. Fungsi DeleteRear() digunakan untuk menghapus elemen pada bagian belakang.
//    - Contoh: [Tino Rino Juno Septo Dono Gio] => [Tino Rino Juno Septo Dono nil]
// 4. Fungsi DeleteFront() digunakan untuk menghapus elemen pada bagian depan.
//    - Contoh: [Tino Rino Juno Septo Dono Gio] => [nil Rino Juno Septo Dono Gio]
// 5. Fungsi Display() digunakan untuk menampilkan isi queue.

package main

import (
	"fmt"
)

// r - rear
// f - front
// n - size of Queue
// q - Queue

type Queue []string

var r, f, n int
var q Queue

func init() {
	r = -1
	f = -1
	n = 4
	q = make(Queue, 5)
	for i := range q {
		q[i] = "nil"
	}
}

func main() {
	for {
		fmt.Printf("\x1bc") // clear screen
		display()
		fmt.Println("Main Menu Queue:")
		fmt.Println("1. INSERT REAR")
		fmt.Println("2. INSERT FRONT")
		fmt.Println("3. DELETE REAR")
		fmt.Println("4. DELETE FRONT")
		fmt.Println("5. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			var rearValue string
			fmt.Print("Enter rear value: ")
			fmt.Scanf("%s\n", &rearValue)
			q, err := q.InsertRear(rearValue)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(q)

			pause()
		case 2:
			var frontValue string
			fmt.Print("Enter front value: ")
			fmt.Scanf("%s\n", &frontValue)
			q, err := q.InsertFront(frontValue)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(q)

			pause()
		case 3:
			y, err := q.deleteRear()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Deleted: ", y)

			pause()
		case 4:
			y, err := q.deleteFront()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Deleted: ", y)

			pause()
		case 5:
			return
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (q Queue) InsertRear(rearValue string) (Queue, error) {
	return q, nil // TODO: replace this
}

func (q Queue) InsertFront(frontValue string) (Queue, error) {
	return q, nil // TODO: replace this
}

func (q Queue) deleteRear() (string, error) {
	if r == -1 {
		return "", fmt.Errorf("element cannot be deleted")
	}
	y := q[r]
	q[r] = "nil"
	if f == r {
		f = -1
		r = -1
	} else {
		r--
	}
	return y, nil
}

func (q Queue) deleteFront() (string, error) {
	if f == -1 {
		return "", fmt.Errorf("underflow")
	}

	y := q[f]
	q[f] = "nil"

	if f == r {
		f = -1
		r = -1
	} else {
		f++
	}
	return y, nil
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
