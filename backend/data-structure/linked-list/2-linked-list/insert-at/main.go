package main

import "fmt"

// Node mewakili node dari linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList mewakili linked list
type LinkedList struct {
	head *Node
	len  int
}

func main() {
	list := new(LinkedList)
	list.head = nil
	list.len = 0

	fmt.Println("Inserting nodes as position...")
	list.InsertAt(0, 12)
	list.InsertAt(-1, 13)
	list.InsertAt(1, 14)
	list.InsertAt(2, 15)

	fmt.Println("Get at node at position:")
	fmt.Println("Node at position 0: ", list.GetAt(0))
	fmt.Println("Node at position -1: ", list.GetAt(-1))
	fmt.Println("Node at position 1: ", list.GetAt(1))
	fmt.Println("Node at position 2: ", list.GetAt(2))
	fmt.Println("Node at position 100: ", list.GetAt(100))
	fmt.Println("Node at position 200: ", list.GetAt(200))
}

// InsertAt memasukkan node baru di posisi yang diberikan dari linked list
func (l *LinkedList) InsertAt(pos int, value int) {
	// membuat node baru
	newNode := Node{}
	newNode.value = value
	// validasi posisi
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

// GetAt mengembalikan node pada posisi yang diberikan dari linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}
