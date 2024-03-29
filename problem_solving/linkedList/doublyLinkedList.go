package linkedList

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type DoublyLinkedList struct {
	head *node
	tail *node
}

func (dl *DoublyLinkedList) Insert(data int) {
	newnode := &node{
		data: data,
		prev: nil,
		next: nil,
	}

	if dl.head == nil {
		dl.head = newnode
		dl.tail = newnode
	} else {
		newnode.prev = dl.tail
		dl.tail.next = newnode
		dl.tail = newnode
	}

}

func (dl *DoublyLinkedList) DisplayFromHead() {
	current := dl.head
	if current == nil {
		fmt.Println("DoublyLinkedlist is empty")
	}
	fmt.Println("DoublyLinked List from head :")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (dl *DoublyLinkedList) DisplayFromTail() {
	current := dl.tail
	if current == nil {
		fmt.Println("DoublyLinkedlist is empty")
	}
	fmt.Println("DoublyLinkedList from tail :")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
}
