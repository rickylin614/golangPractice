package main

import (
	"fmt"
)

//別名的方式改寫
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func main() {
	q := Queue{1}
	q.Push(2)
	q.Push(5)
	fmt.Println(q)
	q.Push(7)
	q.Push(99)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	//先進先出
}
