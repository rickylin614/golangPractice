package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	consts()
	enums()
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4                                    //不定義類型 在使用時 自動給float/int 根據需求
	var c int = int(math.Sqrt(a*a + b*b))                //這邊時a給float
	fmt.Println(filename, c, reflect.TypeOf(a).String()) //反射時 給int
}

func enums() {
	const (
		java = iota // iota 自增值 往下自動  0 1 2 3
		python
		cpp
		golang
	)
	fmt.Println(java, python, cpp, golang) //0 1 2 3
	const (
		b = 1 << (10 * iota) // iota 自增值 往下自動 0 1 2 3 4
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb) // 10^0 10^1
}
