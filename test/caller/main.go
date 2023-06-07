package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		doSomething()
	}()
	time.Sleep(time.Millisecond * 100)
}

func doSomething() {
	functionName := getFunctionName()
	fmt.Println("The function name is:", functionName)
}

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}
