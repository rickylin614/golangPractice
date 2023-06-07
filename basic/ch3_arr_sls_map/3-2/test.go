package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6])
	fmt.Println(arr[:4])
	fmt.Println(arr[4:])
	fmt.Println(arr[:])
	sli := arr[2:6]
	fmt.Println(sli)
	updateSlice(sli)
	fmt.Println(sli)
	fmt.Println(arr)

	sli = sli[1:]
	fmt.Println(sli)
	fmt.Println(cap(sli)) //可得知雖然sli只有三筆數據 但實際容量為5
	sli = sli[2:5]        //有辦法突破原本的sli讀取到arr的部分
	fmt.Println(sli)
	sli[0] = 333333
	fmt.Println(arr)
}

// slice : ptr給首header len給有資料長度 cap給整個總容量

//此時的slice是對arr的指標 因此變動也都會影響到本身值以及ARR本身值
func updateSlice(s []int) {
	s[0] = 100
}
