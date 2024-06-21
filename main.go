package main

import (
	"daniel-moya/data-structures/list"
	"daniel-moya/data-structures/queue"
	"daniel-moya/data-structures/stack"
	"daniel-moya/data-structures/trees"
)

func main() {
	root := trees.GetSortedTree()
	root.PreOrderTraversePrint()
	root.InOrderTraversePrint()
	root.PostOrderTraversePrint()

	queue := queue.GetSampleQueue()
	queue.PrintQueue()

	stack := stack.GetSampleStack()
	stack.PrintStack()

	list := list.GetSampleLinkedList()
	list.Print()

}
