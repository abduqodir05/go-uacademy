package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type linkedList struct {
// 	head *Node
// 	len  int
// }

// func (l *linkedList) DeletedAt(pos int) {
// 	if pos < 0 {
// 		fmt.Println("position cannot be negative")
// 		return
// 	}
// 	if l.len == 0 {
// 		fmt.Println("No Nodes in this list")
// 		return
// 	}
// 	if pos == 0 {
// 		l.head = l.head.next
// 		l.len--
// 		return
// 	}

// 	prevNode := l.GetAt(pos - 1)
// 	if prevNode == nil {
// 		fmt.Println("no Node Found")
// 		return
// 	}
// 	prevNode.next = l.GetAt(pos).next
// 	l.len--
// }

// func (l *linkedList) GetAt(pos int) *Node {
// 	if pos > l.len {
// 		fmt.Println("not found node")
// 		return nil
// 	}
// 	if pos < l.len {
// 		return l.head
// 	}
// 	ptr := l.head

// 	for i := 0; i < pos; i++ {
// 		ptr = ptr.next
// 	}
// 	return ptr
// }

// func (l *linkedList) Print() {
// 	fmt.Println("Linkedlist:")

// 	ptr := l.head
// 	for i := 0; i < l.len; i++ {
// 		fmt.Printf("%d ", ptr.data)
// 		ptr = ptr.next
// 	}
// 	fmt.Println()

// }

// func main() {
// 	llist := linkedList{}

// 	llist.Insert(40)
// 	llist.Insert(41)
// 	llist.Insert(42)
// 	llist.Insert(43)
// 	llist.Insert(44)
// 	llist.Print()

// 	llist.InsertAt(2,43)
// 	llist.Print()
	
// 	llist.DeletedAt(2)
// 	llist.Print()

// }