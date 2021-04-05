package main // go語言定義一樣需要package   package main表示的則是程式進入點

import (
	"fmt"
)

func main() {
	fmt.Print("Hello World") // 1. go 語言console out msg 的方式
	// Print P開頭大寫表示這是Public函數(go語言的function看開頭的大小寫決定是public or private)

	// array
	arr := [3]int{1, 2}       //自定義長度 超出的長度會為0
	arr2 := [...]int{1, 2, 3} //自動長度 數字三所以該陣列長度為3
	fmt.Println(arr)
	fmt.Println(arr2)

	// slice
	sli := []int{}            //不定義長度 即為slice
	sli2 := make([]int, 10)   //使用make自定義長度 int補0
	sli3 := make([]string, 3) //使用string自定義長度 string補空字串
	sli4 := sli2[4:5]         //使用上下標複製值 start -> end - 1 ( 4 -> 4)
	sli2[4] = 4
	sli3[1] = "a"
	sli3[2] = "b"
	fmt.Println(sli)
	fmt.Println(sli2[2:5])
	fmt.Println(sli3)
	fmt.Println(sli4)
	fmt.Println(len(sli2))
	fmt.Println(len(sli2))

}

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello world!")
// }

// func main() {
// 	http.HandleFunc("/", hello)
// 	http.ListenAndServe(":8000", nil)
// }
