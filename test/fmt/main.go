package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var arr [3]string = [3]string{"a", "b", "c"}
	var sli []string = arr[1:]
	arr[2] = "d"
	fmt.Println(sli)

	var a string = "中文字"
	fmt.Printf("%X\n", a)

	x := hex.EncodeToString([]byte(a))
	x = strings.ToUpper(x)
	fmt.Println(x)

	m := make(map[string]interface{}, 0)
	m["aa"] = "aa"
	if v, ok := m["aa"].(int); ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println(v, ok)
	}

}
