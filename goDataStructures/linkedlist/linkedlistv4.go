package main

import "fmt"

type Nodev4 struct {
	data int32
	next *Nodev4
}

type LinkedListv4 struct {
	head   *Nodev4
	length int
}

func (l *LinkedListv4) prepend(n *Nodev4) {
	curr := l.head
	n.next = curr
	l.head = n
}

func (l *LinkedListv4) reverseList() *Nodev4 {
	curr := l.head
	var prev *Nodev4
	var following = l.head

	for curr != nil {
		following = following.next
		curr.next = prev
		prev = curr
		curr = following
	}

	return prev
}

func (l *LinkedListv4) printLinkedList() {
	curr := l.head
	fmt.Println("Printing linked list: ")
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.next
	}
	fmt.Print(" NULL ")
}

func main() {

	llist := LinkedListv4{}

	node1 := &Nodev4{data: 8}
	node2 := &Nodev4{data: 16}
	node3 := &Nodev4{data: 24}

	llist.prepend(node1)

	llist.prepend(node2)

	llist.prepend(node3)

	llist.printLinkedList()

}
