package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f float64 = 1.1
	s := strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Println(s)
}
