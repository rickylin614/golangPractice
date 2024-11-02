package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	s1 := []int{1, 3, 2}
	OriginSort(s1)
	fmt.Println(s1) //

	s2 := []int{1, 3, 2}
	NewSort(s2)
	fmt.Println(s2)

}

func OriginSort(s []int) {
	sort.Ints(s)
}

func NewSort(s []int) {
	slices.Sort(s)
}

type TestData struct {
	// 在這裡添加 TestData 的字段，例如：
	Value int
}

func OriginSort2(s []TestData) {
	// 使用 sort.SliceStable 函數進行結構體切片的排序，這裡可以定義排序規則
	sort.Slice(s, func(i, j int) bool {
		// 在這裡定義你的排序規則，例如按照 Value 字段的值升序排序
		return s[i].Value < s[j].Value
	})
}

func NewSort2(s []TestData) {
	// 使用 sort.SliceStable 函數進行結構體切片的排序，這裡可以定義排序規則
	slices.SortFunc(s, func(a, b TestData) int {
		if a.Value < b.Value {
			return -1
		}
		if a.Value > b.Value {
			return +1
		}
		return 0
	})
}
