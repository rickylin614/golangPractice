package main

import "fmt"

//函數本身可以當成一種變數，傳遞給他人

//參數，變量，返回值都能是一種函數

//函數 -> 必包

func adder() func(int) int {
	// sum := 0
	// return func(value int) int {
	// 	sum += value
	// 	return sum
	// }

	fmt.Println("entry adder()")
	sum := 0
	funcModel := func(value int) int { //函數可直接定義為一個變量 並且裡面可包含外部域的變數
		sum += value
		return sum //返回時，實際上返回的是此FUNC以及自由變量SUM
	}
	fmt.Println("exit adder()")
	return funcModel
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	// return func(v int) (int, iAdder) {
	// 	return base + v, adder2(base + v)
	// }
	funcModel := func(v int) (int, iAdder) {
		sum := base + v
		return sum, adder2(sum)
	}
	return funcModel
}

func main() {
	a := adder() // a 是一個adder內的返回參數
	for i := 0; i < 11; i++ {
		fmt.Println(a(i)) // 調用adder內的func ， 因為都是adder內的func 所以並不會執行到 sum := 0
	}

	b := adder2(0)
	for i := 0; i < 11; i++ {
		var s int
		s, b = b(i)
		fmt.Println(s)
	}

	// b := adder2(10)
	// var c int
	// c, b = b(2)
	// fmt.Println(c)
	// c, _ = b(3)
	// fmt.Println(c)
}
