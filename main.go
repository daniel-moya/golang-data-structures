package main

import (
	"daniel-moya/data-structures/queue"
	"daniel-moya/data-structures/trees"
)

func main() {
	root := trees.GetSortedTree()
	root.PreOrderTraversePrint()
	root.InOrderTraversePrint()
	root.PostOrderTraversePrint()

	queue := queue.GetSampleQueue()
	queue.PrintQueue()
}
