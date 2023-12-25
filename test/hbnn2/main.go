package main

import (
	"fmt"
	"math/rand"
)

const r = 0.5

func main() {
	m := make(map[bool]int, 2)

	for i := 0; i < 100000; i++ {
		b := rand.Float32()
		if r < b {
			m[true]++
		} else {
			m[false]++
		}
	}
	fmt.Println(m)

}
