package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

//function定義: func (接收者) 方法名稱(引數值) (回傳值)
// (int a) add() 寫成 a.add()
// add(int a) 寫成 add(a)
//遍歷treeNode
func (root *TreeNode) Traverse() {
	if root == nil {
		return
	}

	root.Left.Traverse() //左節點有值 就列出左節點
	fmt.Printf("%d ", root.Value)
	root.Right.Traverse() //而後再依序列出右邊的值
}

type AddedValue struct {
	Value int
}
