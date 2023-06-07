package main

import (
	"fmt"
	"time"
)

func main() {
	// any在泛型中，可符合comparable。(一般結構體需要繼承Compare才符合)

	a := 0
	b := time.Now()
	// c := 1.1
	// d := []byte{'a', 'b', 'c'}

	data := map[any]any{ // any does not implement comparable (before 1.20)
		a: b,
	}

	ks := keys(data)
	fmt.Printf("%+v", ks)
}

func keys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SliceToArrayExample() {

}

type Entity struct {
	A string
	B int
}

func (Entity) Compare() {

}
