package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Hour.Nanoseconds())
	for i := 1; i < 10; i++ {
		GetCounts(float64(i))
	}
}

func GetCounts(s float64) map[int]int {
	split := float64(100) / s
	count := make(map[int]int, 0)
	// 賠率 1.67
	for i := 0; i < 100000; i++ {
		data := rand.Intn(int(split)) // rand.Intn(16)  map[0:12383 1:12542 2:12465 3:12535 4:12633 5:12542 6:6211 7:6240 8:6233 9:6216]
		r := data % 10
		count[r]++
	}
	fmt.Printf("%d :count%v\n", int(split), count)
	return count
}
