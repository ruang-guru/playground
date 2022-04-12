package main

import (
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
	list.Insert(20)

	linkedlist, err := list.DeleteVal(14)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Printing nodes")
	for i := 0; i < linkedlist.len; i++ {
		fmt.Println(linkedlist.GetAt(i).value)
	}
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

// DeleteVal menghapus node dengan nilai yang diberikan dari linked list
func (l *LinkedList) DeleteVal(val int) (LinkedList, error) {
	ptr := l.head
	if l.len == 0 {
		return LinkedList{}, fmt.Errorf("list is empty")
	}
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			// TODO: answer here
			return *l, nil
		}
		ptr = ptr.next
	}
	return LinkedList{}, fmt.Errorf("node not found")
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
