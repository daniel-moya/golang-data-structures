package queue

import "fmt"

type Node struct {
	data int
	prev *Node
}

type Queue struct {
	Tail *Node
	Head *Node
}

func (queue *Queue) PrintQueue() {
	fmt.Println("Queue Ordered Print")
	itNode := queue.Head
	for itNode != nil {
		fmt.Println(itNode.data)
		itNode = itNode.prev
	}
}

func (queue *Queue) Pop() {
	queue.Head = queue.Head.prev
}

func (queue *Queue) Push(data int) {
	oldTail := queue.Tail
	queue.Tail = &Node{data: data, prev: nil}
	oldTail.prev = queue.Tail
}

func (queue *Queue) Get() *Node {
	return queue.Head
}

func GetSampleQueue() *Queue {
	seven := &Node{data: 7, prev: nil}
	six := &Node{data: 6, prev: seven}
	five := &Node{data: 5, prev: six}
	four := &Node{data: 4, prev: five}
	three := &Node{data: 3, prev: four}
	two := &Node{data: 2, prev: three}
	one := &Node{data: 1, prev: two}
	queue := &Queue{Head: one, Tail: seven}
	queue.Pop()
	queue.Pop()
	queue.Push(8)
	return queue
}
