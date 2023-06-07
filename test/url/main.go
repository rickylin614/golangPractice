package main

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

type QueryOptions struct {
	Size   int      `url:"size,omitempty"`
	Filter string   `url:"filter,omitempty"`
	Array  []string `url:"array"`
}

func main() {
	q := QueryOptions{
		Size:   1,
		Filter: "abc",
		Array:  []string{"1", "2", "3", "4"},
	}
	vals, _ := query.Values(q)
	fmt.Println(vals)
	fmt.Println(vals.Encode())
}
