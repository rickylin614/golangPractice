package main

import (
	"fmt"
)

func main() {
	var strs []string = []string{"1", "2", "3"}
	var p_strs []*string = []*string{}

	for _, v := range strs {
		p_strs = append(p_strs, &v)
	}
	fmt.Println(p_strs)

	fmt.Println(ABC(1))
}

func ABC(x int) (y int) {
	defer func() {
		if y == 0 {
			y = 1
		} else {
			y++
		}
	}()

	return x + 1
}
