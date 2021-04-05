package main

import "fmt"

func main() {
	// for 條件不須括號 可省略初始 結束 遞增
	sum := 0
	for { //沒有條件式的情況下 可以直接當成while(true)
		sum++
		fmt.Print(sum, " ")
		if sum > 99 {
			fmt.Printf("\n")
			break
		}
	}

	str := "go語言"
	for i, v := range str {
		fmt.Printf("%d %c\n", i, v)
	}
}
