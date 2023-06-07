package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "2W-020|你好啊"
	s2 := "500-00-00M|你好啊"

	a, b := apiErrStringHandler(s)
	c, d := apiErrStringHandler(s2)

	fmt.Println(a, "-", b)
	fmt.Println(c, "-", d)
}

func apiErrStringHandler(str string) (string, string) {
	if len(str) < 1 {
		return "", str
	}

	rep := strings.Replace(str, "\"", "", -1)
	tmp := Split(rep, "|", true)

	if len(tmp) < 2 {
		return "", str
	}

	code := Split(tmp[0], "-", true)

	if len(code) < 3 {

		if tmp[0] == "00-001" || tmp[0] == "00-002" {
			return "401-" + tmp[0], tmp[1]
		}

		if tmp[0] == "00-008" {
			return "400-" + tmp[0], tmp[1]
		}

	}

	return tmp[0], tmp[1]
}

func Split(s, sep string, removeEmptyItem bool) []string {
	items := make([]string, 0, 10)
	for _, v := range strings.Split(s, sep) {
		v = strings.TrimSpace(v)
		if removeEmptyItem {
			if HasValue(v) {
				items = append(items, v)
			}
		} else {
			items = append(items, v)
		}
	}

	return items
}

func HasValue(s string) bool {
	return len(strings.Trim(s, " ")) > 0
}
