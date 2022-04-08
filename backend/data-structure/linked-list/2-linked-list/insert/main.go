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

	fmt.Println("Inserting nodes...")
	list.Insert(12)
	list.Insert(13)
	list.Insert(14)
	list.Insert(15)
	fmt.Println("Printing nodes")
	list.Print()

	fmt.Println("Position of 12 value: ", list.Search(12))
	fmt.Println("Position of 14 value: ", list.Search(14))
	fmt.Println("Position of 15 value: ", list.Search(15))
	fmt.Println("Position of 100 value: ", list.Search(100))
	fmt.Println("Position of 200 value: ", list.Search(200))

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

//  Search mencari posisi dari nilai yang diberikan dari linked list
func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}
