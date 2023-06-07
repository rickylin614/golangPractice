package main

import (
	"fmt"
)

//停止當前函數執行
//一直向上返回 並執行defer
//沒有遇見recover 程序退出

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}() //匿名函數宣告 後加入()表示直接調用
	// tryRecover()
	// panic(errors.New("this is an error"))
	b := 0
	a := 13245678 + (5 / b)
	fmt.Println(a, b)
}

func main() {
	tryRecover()
	fmt.Println("main end")
}
