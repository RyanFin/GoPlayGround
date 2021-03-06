package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	// old head assigned as the second node
	second := l.head
	// assign new node 'n' to the old head
	l.head = n
	// attach current head to the newly created second element
	l.head.next = second
	// increase the overall length of the linkedlist
	l.length++
}

func (l LinkedList) printOutData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d\n", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

// insert data into LinkedList given a specified position in the list
// The head node of the linkedlist marks the first parameter
func (l *LinkedList) InsertNode(head *Node, data int, pos int) {
	// generate a new node for insertion
	newNode := &Node{data: data}

	// set current Node to the head node n
	currentNode := head

	// create a new index
	var index = 0

	// while loop...
	for index < pos-1 {
		currentNode = currentNode.next
		index++
	}

	newNode.next = currentNode.next
	currentNode.next = newNode

}

func main() {
	// create an empty linked list
	myList := LinkedList{}

	//create a node with some data
	node1 := &Node{data: 71} // we need the memory address for this newly created node for referal purposes
	node2 := &Node{data: 54}
	node3 := &Node{data: 91}
	node4 := &Node{data: 88}
	node5 := &Node{data: 12}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)

	myList.InsertNode(myList.head, 1997, 0)
	myList.printOutData()

}
