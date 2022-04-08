// Tulis sebuah fungsi untuk menghapus sebuah node dalam sebuah single-linked list.
// Kita tidak akan diberikan akses ke head dari list, melainkan kita akan diberikan akses ke node yang akan dihapus secara langsung.
// Dapat dipastikan node yang akan dihapus bukan merupakan tail node dalam list.
//
// Contoh 1
// Input: head = [4,5,1,9], node = 5
// Output: [4,1,9]
// Explanation: Kita mengetahui node 5 akan dihapus, sehingga linked list menjadi 4 -> 1 -> 9 setelah memanggil fungsi delete.
//
// Contoh 2
// Input: head = [4,5,1,9], node = 1
// Output: [4,5,9]
// Explanation: Kita mengetahui node 5 akan dihapus, sehingga linked list menjadi 4 -> 5 -> 9 setelah memanggil fungsi delete.

package main

import "fmt"

// Definisi untuk singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := new(ListNode)
	node1.Val = 4

	node2 := new(ListNode)
	node2.Val = 5
	node1.Next = node2

	node3 := new(ListNode)
	node3.Val = 1
	node2.Next = node3

	node4 := new(ListNode)
	node4.Val = 9
	node3.Next = node4

	deleteNode1(node2)

	display(node2)
}

func deleteNode1(node *ListNode) {
	affectedNode := node.Next
	node.Val = affectedNode.Val
	node.Next = affectedNode.Next
	affectedNode = nil
}

func display(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
