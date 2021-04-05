package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "每日負能量FUCK"
	for i, v := range s {
		fmt.Printf("index: %d , value: %c , unicode : %X\n", i, v, v)
	}
	fmt.Println()
	for i, v := range []byte(s) {
		fmt.Printf("index: %d , value: %c , unicode : %X\n", i, v, v)
	}

	fmt.Println()
	for i, v := range []rune(s) {
		fmt.Printf("index: %d , value: %c , unicode : %X\n", i, v, v)
	}
	// utf8.RuneCountInString(s) = 9 //直接取得字符數
	ch, size := utf8.DecodeRune([]byte(s)) //計算byte第一位
	fmt.Printf("%x %c %d", ch, ch, size)

	//常用字符串操作 : Fields,Split,Join
	//Contains,Index
	//ToLower,ToUpper
	//Trim,TrimRight,TrimLeft
}
