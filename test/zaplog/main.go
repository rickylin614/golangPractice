package main

import (
	"errors"

	"go.uber.org/zap"
)

var zaplog *zap.Logger

func main() {
	zaplog = zap.NewExample(
		// zap.AddCaller(),
		zap.AddStacktrace(zap.InfoLevel),
	)

	defer func() {
		if err := recover(); err != nil {
			zaplog.Error("ERROR!", zap.Error(err.(error)))
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
	// panic(errors.New("err!!"))
	zaplog.Error("msg", zap.Error(errors.New("connect time out")))
}
