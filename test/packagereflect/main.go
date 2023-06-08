package main

import (
	"fmt"
	"go/types"

	_ "practice/test/packagereflect/modeltest"
)

func main() {

	// create a new type checker
	var conf types.Config
	pkg, err := conf.Check("/test/packagereflect/modeltest", nil, nil, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(pkg)

	// get all types in the package
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		fmt.Println(obj)
		// if obj.Kind() == types.TypeName {
		// 	typ := obj.Type()
		// 	if s, ok := typ.Underlying().(*types.Struct); ok {
		// 		// print the struct name and fields
		// 		fmt.Printf("Struct Name: %s\n", name)
		// 		for i := 0; i < s.NumFields(); i++ {
		// 			field := s.Field(i)
		// 			fmt.Printf("\tField %d: %s %s\n", i+1, field.Name(), field.Type())
		// 		}
		// 	}
		// }
	}

	pkg = types.NewPackage("github.com/rickylin614/common/constants", "name")
	scope = pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		fmt.Println(obj)
	}
}
