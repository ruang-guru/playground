// Diberikan head dari singly linked list, kembalikan middle node dari inked list.
// Jika ada dua middle node, kembalikan middle node kedua.
//
// Contoh 1
// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: middle node nya 3, maka dikembalikan nilai 3 beserta dengan next value nya, jadi 3, 4, 5
//
// Contoh 2
// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: middle node nya 3 dan 4, kembalikan middle node kedua, maka dikembalikan nilai 4 beserta dengan next value nya, jadi 4, 5, 6

package main

import "fmt"

// Definisi untuk singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := new(ListNode)
	node1.Val = 1

	node2 := new(ListNode)
	node2.Val = 2
	node1.Next = node2

	node3 := new(ListNode)
	node3.Val = 3
	node2.Next = node3

	node4 := new(ListNode)
	node4.Val = 4
	node3.Next = node4

	node5 := new(ListNode)
	node5.Val = 5
	node4.Next = node5

	data := node1.MiddleNode(node1)

	// Print Middle Node
	fmt.Println("Middle Node: ", data)
	// Print Next Node
	fmt.Println("Next Node: ", data.Next)

	// Print List
	fmt.Println("List: ")
	for data != nil {
		fmt.Println(data.Val)
		data = data.Next
	}

}

func (node ListNode) MiddleNode(head *ListNode) *ListNode {
	l := 0
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		l++
	}
	// TODO: answer here
	return head
}
