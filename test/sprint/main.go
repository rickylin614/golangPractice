package main

import (
	"fmt"
	"time"
)

type mockEnum int

func (mockEnum) String() string {
	return "Mock"
}

const (
	mockA mockEnum = iota
	mockB
)

func main() {
	time.Local, _ = time.LoadLocation("Europe/Malta")
	t := time.Now()
	fmt.Println(t)

	var A interface{} = mockA
	a := fmt.Sprint(A)
	fmt.Println(a)
}
