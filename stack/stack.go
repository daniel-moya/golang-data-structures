package stack

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	Head *Node
}

func (stack *Stack) PrintStack() {
	fmt.Println("Stack Ordered Print")
	itNode := stack.Head
	for itNode != nil {
		fmt.Println(itNode.data)
		itNode = itNode.next
	}
}

func (stack *Stack) Pop() {
	stack.Head = stack.Head.next
}

func (stack *Stack) Push(data int) {
	stack.Head = &Node{data: data, next: stack.Head}
}

func (stack *Stack) Get() *Node {
	return stack.Head
}

func GetSampleStack() *Stack {
	one := &Node{data: 1, next: nil}
	two := &Node{data: 2, next: one}
	three := &Node{data: 3, next: two}
	four := &Node{data: 4, next: three}
	five := &Node{data: 5, next: four}
	six := &Node{data: 6, next: five}
	seven := &Node{data: 7, next: six}
	stack := &Stack{Head: seven}
	stack.Pop()
	stack.Pop()
	stack.Push(8)
	return stack
}
