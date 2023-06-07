package Queue

import "fmt"

//將會自動輸出文檔 以及可以直接做為測試用
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(5)
	fmt.Println(q.Pop())
	q.Push(7)
	q.Push(99)
	fmt.Println(q.Pop())
	q.Pop()
	fmt.Println(q.Pop())

	// Output:
	// 1
	// 2
	// 7
}
