package main

import "fmt"

func main() {
	var b Z = B{A: A{}}
	b.Do1()
}

type Z interface {
	Do1()
	Do2()
}

type A struct {
	Z
}

func (a A) Do1() {
	a.Z.Do2()
}

func (A) Do2() {
	fmt.Println("A Do2")
}

type B struct {
	A
}

func (B) Do2() {
	fmt.Println("B Do2")
}
