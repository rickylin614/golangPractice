package main

import "fmt"

func main() {
	var root TreeNode
	root.Value = 123
	fmt.Println(root.Value, root.Left, root.Right)
	root = TreeNode{Value: 3}                               //結構體的創建方法
	root.Left = &TreeNode{}                                 //結構體的創建方法
	root.Right = &TreeNode{Value: 5, Left: nil, Right: nil} //結構體的創建方法
	root.Right.Left = &TreeNode{Value: 4}                   //結構體的創建方法
	root.Left.Right = &TreeNode{Value: 2}                   //結構體的創建方法
	root.Traverse()

	fmt.Println()

	c := root.TraverseWithChaneel()
	maxNode := 0
	for node := range c {
		// node <- c 將channel接收到的值給node
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println(maxNode)
}

func (node *TreeNode) TraverseWithChaneel() chan *TreeNode {
	nodeChan := make(chan *TreeNode)
	go func() { //建立一個goroutine 將node值給予channel
		node.TraverseFunc(func(node *TreeNode) {
		})
		nodeChan <- node //遍歷將所有節點給予nodeChan 接收給main:20
		close(nodeChan)
	}()
	return nodeChan
}

func (root *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if root == nil {
		return
	}
	root.Left.TraverseFunc(f)
	f(root)
	root.Right.TraverseFunc(f)
}

func (root *TreeNode) Traverse() {
	if root == nil {
		return
	}

	root.Left.Traverse() //左節點有值 就列出左節點
	fmt.Printf("%d ", root.Value)
	root.Right.Traverse() //而後再依序列出右邊的值
}

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}
