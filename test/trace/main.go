package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Trace Method1 dubug.Stack", string(debug.Stack()))
			// errors.
			// err, _ := err.(errors.StackTrace)
			// fmt.Println("Trace Method2 errors.StackTrace")
		}
	}()
	A()
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	D()
}

var count = 0

func D() {
	panic(errors.New("err!!"))
}
