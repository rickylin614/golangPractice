package main

import (
	"fmt"
	"reflect"
)

type IA interface {
	Say()
}

type IB interface {
	Walk()
}

type Real struct {
	A IA
	B IB
}

var ActionMap map[reflect.Type]interface{}

func init() {
	// var a IA
	// var b IB

	// ActionMap = map[reflect.Type]interface{}{
	// 	IA: IA,
	// }
}

func main() {

}

func Switch(data interface{}) {
	r := Real{}

	var a IA
	fmt.Println(reflect.TypeOf(a).Elem())

	switch x := data.(type) {
	case IA:
		r.A = x
	case IB:
		r.B = x
	}

	structType := reflect.TypeOf(r)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	for i := 0; i < structType.NumField(); i++ {
		// structType.Field(i).Type
		switch x := data.(type) {
		default:
			fmt.Println(x)
			// structType.Field(i).Type ==
		}
	}

	fmt.Println(r)
}
