package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//go語言的函數宣告返回值寫後面
//可返回多個值 函覆本身可作為參數 沒有默認參數 可選參數
func eval(a, b int, op string) (r int, e error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		// panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
	// q = a / b //可以這樣直接定義回傳 但不建議這樣使用 代碼寫的長 有時候會誤判斷
	// r = a % b
	// return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name() //根據傳進來的op位址(函數的位址) 去調用並且給參數
	fmt.Printf("Call Function %s with args %d %d\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(eval(5, 10, "*"))
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3, 4))
	//匿名函數
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
}
