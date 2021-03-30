// Hackerrank solution to delete some data at a specific point in a linkedlist
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
    head *SinglyLinkedListNode
    tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
    node := &SinglyLinkedListNode {
        next: nil,
        data: nodeData,
    }

    if singlyLinkedList.head == nil {
        singlyLinkedList.head = node
    } else {
        singlyLinkedList.tail.next = node
    }

    singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
    for node != nil {
        fmt.Fprintf(writer, "%d", node.data)

        node = node.next

        if node != nil {
            fmt.Fprintf(writer, sep)
        }
    }
}

// Complete the deleteNode function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func deleteNode(head *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
    curr := head
    prev := head 
    index := 0
    
    // deleting the head node
    if position == 0{
        curr = head
        head = head.next
        return head
    }

    for index < int(position){
        prev = curr
        curr = curr.next
        index++
    }
    
    prev.next = curr.next
    curr.next = nil
    
    fmt.Println("index: " , 0)
    fmt.Println("Prev", prev)
    fmt.Println("Curr: " , curr)
    

    return head
     
}

func main() {