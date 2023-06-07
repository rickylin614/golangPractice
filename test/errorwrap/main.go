package main

import (
	"demo/apmdemo/zlog"
	// "errors"
	"errors"
	"fmt"
	// "golang.org/x/errors"
)

func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func main() {
	err := C()
	for err != nil {
		fmt.Println(err)
		err = errors.Unwrap(err)
	}
}

func A() error {
	err := errors.New("A() new error")
	zlog.Error(err)
	return err
}

func B() error {
	if err := A(); err != nil {
		return fmt.Errorf("B() %w", err) // 業務邏輯 要給錢端看的東西
	}
	return nil
}

func C() error {
	if err := B(); err != nil {
		return fmt.Errorf("C() %w", err)
	}
	return nil
}
