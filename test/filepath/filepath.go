package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("./213")
	fmt.Printf("path: %s, err: %v\n", path, err)
	// fmt.Printf("%v %v\n", path, err)
}
