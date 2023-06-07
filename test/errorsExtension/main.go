package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

}

type errorExtension struct {
	StatusCode string
	ErrorMsg   string
	Err        error
	Parameter  []interface{}
}

func (m *errorExtension) Error() string {
	return m.Err.Error()
}

func newErrorExtension(code, format string, args ...interface{}) error {
	var errExt errorExtension
	if len(args) > 0 {
		errExt = errorExtension{

			StatusCode: strings.ToLower(code),
			Err:        errors.New(code + "|" + fmt.Sprintf(format, args...)),
			ErrorMsg:   fmt.Sprintf(format, args...),
			Parameter:  args,
		}
	} else {
		errExt = errorExtension{
			StatusCode: strings.ToLower(code),
			Err:        errors.New(code + "|" + format),
			ErrorMsg:   format,
		}
	}
	return &errExt
}
