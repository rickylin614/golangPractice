package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Trace Method1 dubug.Stack", err, string(debug.Stack()))
		}
		fmt.Println("12333")
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

	i, err := rand.Int(rand.Reader, big.NewInt(int64(0)))
	fmt.Println(i, err)

	// goroutine的panic會強制程序結束且無法被main的defer接到。需要謹慎使用。
	// go func () {
	// 	panic(errors.New("err!!"))
	// }()
}
