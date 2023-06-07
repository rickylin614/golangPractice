package main

import "fmt"

func main() {
	//arrays
	var arr0 [5]int
	arr1 := [3]int{1, 2}      //自定義長度 超出的長度會為0
	arr2 := [...]int{1, 2, 3} //自動長度 數字三所以該陣列長度為3
	arr3 := [4][5]int{}
	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	sum := 0
	for _, v := range arr2 {
		sum += v
	}
	fmt.Println(sum)
}
