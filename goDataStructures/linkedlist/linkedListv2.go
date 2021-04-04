package main

import "fmt"

type Nodev2 struct {
	data int
	next *Nodev2
}

type LinkedListv2 struct {
	head *Nodev2
	len  int
}

func (l *LinkedListv2) prepend(n *Nodev2) {
	second := l.head
	l.head = n
	l.head.next = second
	l.len++
}

func (l LinkedListv2) first() *Nodev2 {
	return l.head
}

// func (n Nodev2) Next() *Nodev2 {
// 	return n.next
// }

func (l LinkedListv2) printAllData() {
	toPrint := l.head
	for l.len != 0 {
		// fmt.Printf("%d ", toPrint)
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.len--
	}
	fmt.Println(" ")
}

func (n *Nodev2) insertNodeAtHead(l LinkedListv2, dat int) {
	node := &Nodev2{data: dat}
	node.next = l.head
}

func (l *LinkedListv2) removeElemWithPos(pos int) *LinkedListv2 {
	curr := l.head
	prev := l.head

	// delete head node if position == 0
	if pos == 0 {
		curr = l.head
		l.head = l.head.next
		return l
	}

	index := 0
	for index < pos {
		prev = curr
		curr = curr.next
		index++
	}

	prev.next = curr.next
	curr.next = nil

	return l

}

func reverse(head *Nodev2) *Nodev2 {
	// set previous to equal nil
	var previous *Nodev2
	current := head
	following := head

	// while curr is not equal to nil (end) iterate to the end of the list
	for current != nil {
		// following is equal to the next node
		following = following.next
		// current points to previous node to redirect node link
		current.next = previous
		// previous is set to current
		previous = current
		// set current to the next node
		current = following
	}

	return previous
}

func main() {
	myList := LinkedListv2{}

	n1 := &Nodev2{data: 45}
	n2 := &Nodev2{data: 87}
	n3 := &Nodev2{data: 100}
	n4 := &Nodev2{data: 156}
	n5 := &Nodev2{data: 71}
	n6 := &Nodev2{data: 12}
	nf := &Nodev2{}

	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.prepend(n4)
	myList.prepend(n5)
	myList.prepend(n6)
	nf.insertNodeAtHead(myList, 202)

	// fmt.Println(myList.first())

	myList.printAllData()
	myList.removeElemWithPos(2)
	// myList.printAllData()

	// currentNode := myList.head

	// fmt.Println(currentNode.Next())
	// fmt.Println(currentNode.Next())

}
