package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	smap := sync.Map{}
	smap.Store("a", "d")
	smap.Store("b", "c")
	smap.Store("c", "b")
	smap.Store("d", "a")

	omap := map[string]any{
		"a": "d",
		"b": "c",
		"c": "b",
		"d": "a",
	}

	s, err := json.Marshal(smap)
	fmt.Printf("%s,%s", s, err)
	s, err = json.Marshal(omap)
	fmt.Printf("%s,%s", s, err)
}
