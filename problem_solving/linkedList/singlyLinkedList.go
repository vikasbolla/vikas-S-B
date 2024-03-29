package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (sl *LinkedList) Insert(data int) {
	newnode := &Node{
		data: data,
	}
	if sl.head == nil {
		sl.head = newnode
	} else {
		current := sl.head
		for current.next != nil {
			current = current.next
		}
		current.next = newnode
	}

}

func AddNodeAtEnd(head *Node, data int) *Node {
	newnode := &Node{data: data, next: nil}
	if head == nil {
		return newnode
	}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newnode
	return head
}

func AddNodeAtBegging(head *Node, data int) *Node {
	newnode := &Node{data: data, next: head}

	return newnode
}

func DeleteNode(head *Node, value int) *Node {
	if head != nil && head.data == value {
		return head.next
	}
	var prevNode *Node
	current := head
	for current != nil && current.data != value {
		prevNode = current
		current = current.next
	}
	if current != nil {
		prevNode.next = current.next
	}
	return head
}

func ArrayToLinkedList(arr []int) *Node {
	var head, tail *Node
	for _, val := range arr {
		newnode := &Node{
			data: val,
			next: nil,
		}
		if head == nil {
			head = newnode
			tail = newnode
		} else {
			tail.next = newnode
			tail = newnode
		}
	}
	return head
}
func PrintConvertedList(head *Node) {
	current := head
	fmt.Println()
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
func (sl *LinkedList) Display() {
	current := sl.head
	if current == nil {
		fmt.Println("LinkedList is empty")
	}
	fmt.Println("LinkedList :")
	for current.next != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func InsertNodeAfter(head *Node, value int, newdata int) *Node { // insertNodeAfter([10,20,30,40], 30, 35)
	currentNode := head
	for currentNode != nil {
		if currentNode.data == value {
			newnode := &Node{
				data: newdata,
				next: currentNode.next,
			}
			currentNode.next = newnode
			break

		}
		currentNode = currentNode.next

	}
	return head
}

func InsertNodeBefore(head *Node, value int, newdata int) *Node { // insertNodeBefore([10,20,30,40], 30, 35)
	if head == nil {
		return head
	}
	if head.data == value {
		newnode := &Node{
			data: newdata,
			next: head,
		}
		return newnode
	}
	currentNode := head
	for currentNode != nil {
		if currentNode.next.data == value {
			newnode := &Node{
				data: newdata,
				next: currentNode.next,
			}
			currentNode.next = newnode
			break
		}
		currentNode = currentNode.next
	}
	return head
}

func InsertNodeBefore2(head *Node, value int, newData int) *Node {
	if head == nil {
		return head
	}

	if head.data == value {
		newNode := &Node{
			data: newData,
			next: head,
		}
		return newNode
	}

	currentNode := head
	var prevNode *Node

	for currentNode != nil {
		if currentNode.data == value {
			newNode := &Node{
				data: newData,
				next: currentNode,
			}
			prevNode.next = newNode
			break
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}

	return head
}
