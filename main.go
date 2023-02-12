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

func (l *LinkedList) Deleted(pos int) {

	if pos < 0 {
		fmt.Println("position can not be negative")
		return
	}

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}

	if pos == 0 {
		l.head = l.head.next
		l.len--
		return
	}

	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		return
	}

	prevNode.next = l.GetAt(pos).next
	l.len--
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

	llist.Insert(43)
	llist.Insert(44)
	llist.Insert(40)
	llist.Insert(42)
	llist.Insert(41)

	llist.Deleted(2)
	llist.Print()
}

// stack vazifasini qlomadim