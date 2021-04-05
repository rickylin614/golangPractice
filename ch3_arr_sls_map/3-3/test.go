package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]       // 2 3 4 5
	s2 := s1[3:5]        // 5 6
	s3 := append(s2, 10) // 5 6 後增加 10
	s4 := append(s3, 11) // 5 6 10後增加 11
	s5 := append(s4, 12) // 5 6 10 11 後增加 12 //append如果超出原本範圍 將創建一個容量更大的arr在底層
	fmt.Println(arr)     //因為長度定義為 8 只記錄 0 1 2 3 4 5 6 10 (11 12超出範圍)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
	printSlice(s4)
	printSlice(s5)

	var s []int // Zero value for slice is nil

	for i := 0; i < 50; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	printSlice(s)

	ss := []int{2, 4, 6, 8, 10}
	ss2 := make([]int, 16) //make(type,len,cap)
	printSlice(ss)
	printSlice(ss2)
	copy(ss2, ss) // 把ss2的前五個元素變為 ss的元素 //只影響值不影響長度 超出不動作
	printSlice(ss2)
	//組合ss 與 ss2
	ss2 = append(ss2, ss[:]...) //用...可以將元素都轉移
	printSlice(ss2)
	//刪除中間的元素
	ss2 = append(ss2[:2], ss2[4:]...) // 2 ~ 3的元素被刪除
	printSlice(ss2)
}

func printSlice(s []int) {
	fmt.Printf("lens=%d,cap=%d", len(s), cap(s))
	fmt.Println(s)
}
