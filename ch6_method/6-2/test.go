package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	f := fibonacci() //前兩者之合

	// f() // 1
	printFileContents(f)
	// f() // 1
	// f() // 2
	// f() // 3
	// f() // 5
	// f() // 8
	// f() // 13
	// f() // 21
	// f() // 34
	// f() // 55

	root := Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 2}
	root.Traverse()
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d ", next)
	n, err = strings.NewReader(s).Read(p)
	return
}

func printFileContents(reader io.Reader) { //io.Reader 為一個inteface 繼承者必須要有Read方法
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("%s \n", scanner.Text())
	}
}

//斐波納契數列
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Traverse() {
	//匿名函數方式定義func的內容 並在呼叫TraverseFunc
	nodeCount := 0
	node.TraverseFunc(func(node *Node) {
		nodeCount++ //直接引進自由變量
		fmt.Println(node.Value)
	})
	fmt.Println("nodeCount:", nodeCount)
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
