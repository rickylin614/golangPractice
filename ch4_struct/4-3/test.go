package main

import (
	"fmt"
	"myGolang/ch4_struct/4-2/tree"
)

//如何擴充系統類型/別人的類型

//1. 定義別名
//2. 使用組合

//宣告一個已有的結構，在此結構另外組合自己需要的方法以及參數
type MyTreeNode struct {
	node *tree.TreeNode
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	fmt.Print(myNode.node.Value, " ")
}

func main() {
	// var root tree.TreeNode
	root := tree.TreeNode{Value: 3}                              //結構體的創建方法
	root.Left = &tree.TreeNode{}                                 //結構體的創建方法
	root.Right = &tree.TreeNode{Value: 5, Left: nil, Right: nil} //結構體的創建方法
	root.Right.Left = new(tree.TreeNode)                         //結構體的創建方法
	root.Left.Right = tree.CreateNode(2)                         //結構體的創建方法
	myRoot := MyTreeNode{&root}
	myRoot.postOrder()
}
