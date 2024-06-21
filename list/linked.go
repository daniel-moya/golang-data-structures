package list

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	first *Node
	last  *Node
}

func (list *LinkedList) Print() {
	fmt.Println("Linked list")
	it := list.first
	for it != nil {
		fmt.Println(it.data)
		it = it.next
	}
}

func (list *LinkedList) Unshift(data int) {
	oldFirst := list.first
	list.first = &Node{data: data, next: oldFirst}
}

func (list *LinkedList) Push(data int) {
	newNode := &Node{data: data}
	if list.first == nil {
		list.first = newNode
		list.last = newNode
		return
	}

	list.last.next = newNode
	list.last = newNode
}

func GetSampleLinkedList() *LinkedList {
	list := &LinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)
	list.Push(6)
	list.Push(7)
	list.Unshift(0)
	return list
}
