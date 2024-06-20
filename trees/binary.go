package trees

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

var four *Node = &Node{Data: 4}
var five *Node = &Node{Data: 5}
var six *Node = &Node{Data: 6}
var seven *Node = &Node{Data: 7}
var two *Node = &Node{Data: 2, Left: four, Right: five}
var three *Node = &Node{Data: 3, Left: six, Right: seven}
var root *Node = &Node{Data: 1, Left: two, Right: three}

func (root *Node) PreOrderTraversePrint() {
	fmt.Println("Depth First Search: Pre Order traverse")
	printPreOrder(root)
}

func (root *Node) InOrderTraversePrint() {
	fmt.Println("Depth First Search: In Order traverse")
	printInOrder(root)
}

func (root *Node) PostOrderTraversePrint() {
	fmt.Println("Depth First Search: Post Order traverse")
	printPostOrder(root)
}

func printPreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Data)
	printPreOrder(root.Left)
	printPreOrder(root.Right)
}

func printInOrder(leftIt *Node) {
	if leftIt == nil {
		return
	}

	printInOrder(leftIt.Left)
	fmt.Println(leftIt.Data)
	printInOrder(leftIt.Right)
}

func printPostOrder(leftIt *Node) {
	if leftIt == nil {
		return
	}

	printPostOrder(leftIt.Left)
	printPostOrder(leftIt.Right)
	fmt.Println(leftIt.Data)

}

func GetSortedTree() *Node {
	return root
}
