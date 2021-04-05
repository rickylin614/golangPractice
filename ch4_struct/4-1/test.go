package main

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

//function定義: func (接收者) 方法名稱(引數值) (回傳值)
// (int a) add() 寫成 a.add()
// add(int a) 寫成 add(a)
//遍歷treeNode
func (root *TreeNode) traverse() {
	if root == nil {
		return
	}

	root.Left.traverse() //左節點有值 就列出左節點
	fmt.Printf("%d ", root.Value)
	root.Right.traverse() //而後再依序列出右邊的值
}

type AddedValue struct {
	Value int
}

func (origin AddedValue) add(a int) (b, c int) {
	b = a + 1
	c = a + 2
	origin.Value = a + b + c
	fmt.Println(origin.Value)
	return
}

func main() {
	var root TreeNode
	fmt.Println(root.Value, root.Left, root.Right)
	root = TreeNode{Value: 3}           //結構體的創建方法
	root.Left = &TreeNode{}             //結構體的創建方法
	root.Right = &TreeNode{5, nil, nil} //結構體的創建方法
	root.Right.Left = new(TreeNode)     //結構體的創建方法
	root.Left.Right = createNode(2)     //結構體的創建方法
	root.traverse()
	fmt.Println()
	value := AddedValue{Value: 100}
	b, c := value.add(9)
	fmt.Println(value, b, c)
}
