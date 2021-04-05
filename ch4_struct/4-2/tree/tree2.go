package tree

import "fmt"

func (origin AddedValue) Add(a int) (b, c int) {
	b = a + 1
	c = a + 2
	origin.Value = a + b + c
	fmt.Println(origin.Value)
	return
}
