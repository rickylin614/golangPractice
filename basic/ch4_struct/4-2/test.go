package main

import (
	"fmt"
	"practice/basic/ch4_struct/4-2/tree"
)

//封裝
// 首字母大寫表示public
// 首字母小寫表示private
// 此特性針對package

//package 包
//每個目錄一個包 同個目錄裡面的package只能有一個名稱
//只能有一個main包
//main包包含可執行入口

func main() {
	var root tree.TreeNode
	root.Value = 123
	fmt.Println(root.Value, root.Left, root.Right)
	root = tree.TreeNode{Value: 3}                               //結構體的創建方法
	root.Left = &tree.TreeNode{}                                 //結構體的創建方法
	root.Right = &tree.TreeNode{Value: 5, Left: nil, Right: nil} //結構體的創建方法
	root.Right.Left = new(tree.TreeNode)                         //結構體的創建方法
	root.Left.Right = tree.CreateNode(2)                         //結構體的創建方法
	root.Traverse()
	// fmt.Println()
	value := tree.AddedValue{Value: 100}
	b, c := value.Add(9)
	fmt.Println(value, b, c)
}
