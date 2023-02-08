package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(data int) {

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

func (l *LinkedList) GetAt(pos int) *Node {

	if pos > l.len {
		fmt.Println("Not found node")
		return nil
	}

	if pos < 0 {
		return l.head
	}

	ptr := l.head

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr
}

func (l *LinkedList) Print() {

	fmt.Println("LinkedList:")

	ptr := l.head

	for i := 0; i < l.len; i++ {

		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}

	fmt.Println()
}

func main() {
	llist := LinkedList{}

	llist.Insert(42)
	llist.Insert(40)
	llist.Insert(44)
	llist.Insert(41)

	llist.Print()

	fmt.Println(llist.GetAt(3))
}
