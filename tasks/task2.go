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

func (l *linkedList) InsertFirst(data int) {
	newNode := Node{}
	newNode.data = data	
	
	if l.len == 0 {
		l.head = &newNode
		l.len++
		return
	}

	newNode.next = l.head
	l.head = &newNode
}
func (l *linkedList) Print() {
	fmt.Println("Linkedlist:")

	ptr := l.head

	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ",ptr.data)
		ptr = ptr.next
	}
	fmt.Println()
}
func main() {
	llist := linkedList{}

	llist.Insert(40)
	llist.Insert(41)
	llist.Insert(42)
	llist.Insert(43)
	llist.Insert(31)
	llist.Insert(44)
	llist.Print()
}