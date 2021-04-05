package main

import (
	"fmt"
)

//指針不能運算
func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	pointerTest(&a)
	fmt.Println(a)
}

func pointerTest(a *int) {
	(*a)++
}
