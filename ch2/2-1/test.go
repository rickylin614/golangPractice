package main

import "fmt"

//外部定義一定要有var 不能:= //變量 作用域範圍for package
var aa = 3
var ss = "kkk"

func main() {
	fmt.Println("HelloWorld")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableAutoDefined()
	fmt.Println(aa, ss)
}

func variableZeroValue() {
	var a int
	var s string //定義完就有零值
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableAutoDefined() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}
