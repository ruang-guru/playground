package main

import (
	"errors"
	"fmt"
)

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

	fmt.Println("Inserting nodes...")
	list.Insert(12)
	list.Insert(13)
	list.Insert(14)
	list.Insert(15)

	fmt.Println("Deleting node at position 2")
	list.DeleteAt(2)
	fmt.Println("Printing nodes")
	list.Print()

}

// Insert memasukkan node baru di akhir dari linked list
func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// DeleteAt menghapus node pada posisi yang diberikan dari linked list
func (l *LinkedList) DeleteAt(pos int) error {
	// validasi posisi
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("no nodes in list")
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos).next
	l.len--
	return nil
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

// Print menampilkan semua node dari linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}
