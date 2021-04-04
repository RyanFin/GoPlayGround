package main

import (
	"errors"
	"fmt"
)

type Nodev3 struct {
	data int
	next *Nodev3
}

type LinkedListv3 struct {
	head   *Nodev3
	length int
}

func (l *LinkedListv3) insert(n *Nodev3) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedListv3) printAllElems() (bool, error) {

	if l.head == nil {
		return false, errors.New("Error printing out list elems....")
	} else {
		curr := l.head
		for curr != nil {
			fmt.Print(curr.data, " -> ")
			curr = curr.next
		}
		fmt.Print("NULL\n")

		return true, nil
	}

	return true, nil
}

// reverse function
func (l *LinkedListv3) reverse() *Nodev3 {
	var previous *Nodev3
	curr := l.head
	following := l.head

	for curr != nil {
		following = following.next

		curr.next = previous

		previous = curr

		curr = following

	}

	return previous
}

func main() {
	llist := LinkedListv3{}

	n1 := &Nodev3{data: 12}
	n2 := &Nodev3{data: 24}
	n3 := &Nodev3{data: 36}

	llist.insert(n1)
	llist.insert(n2)
	llist.insert(n3)

	// print out elems
	_, err := llist.printAllElems()
	if err != nil {
		fmt.Println("Error printing out linked list...")
	}

	llist.reverse()
	_, err = llist.printAllElems()
	if err != nil {
		fmt.Println("Error printing out linked list...")
	}

}
