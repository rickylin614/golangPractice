package main

import (
	"fmt"
	"regexp"
)

func main() {

	// re := regexp.MustCompile(`^[a-zA-Z \p{Han}]+$`)
	re := regexp.MustCompile("^[a-zA-Z \\x{4e00}-\\x{9fff}]+$")
	// re := regexp.MustCompile("([a-zA-Z]+|[\\p{Han}]+|[\\s]+)")
	// 測試字符串
	testString := "a中文字a"
	// 使用正則表達式對字符串進行匹配
	matched := re.MatchString(testString)
	fmt.Println(matched) // 輸出 true
}
