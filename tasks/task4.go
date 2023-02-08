package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head *Node
	len  int
}

func (l *linkedList) Insert(data int) {
	newNode := Node{}
	newNode.data = data

	if l.len == 0 {
		l.head = &newNode
		l.len++
		return
	}
	ptr := l.head

	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *linkedList) GetAt(pos int) *Node {

	if pos > l.len {
		fmt.Println("not found node")
		return nil
	}
	if pos < l.len {
		return l.head
	}
	ptr := l.head

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *linkedList) Print() {
	fmt.Println("Linkedlist:")

	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}
	fmt.Println()

}

func main() {
	
	llist := linkedList{}

	llist.Insert(33)
	llist.Insert(23)
	llist.Insert(53)
	llist.Insert(43)

	fmt.Println(llist.GetAt(1))
	llist.Print()
}