package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	contextFunc(ctx)
	fmt.Println(ctx.Value("key"))
	fmt.Printf("%s\n", ctx)
	ctx2 := context.WithValue(ctx, "key2", "value2")
	fmt.Println(ctx2.Value("key2"))
}

func contextFunc(c context.Context) {
	c = context.WithValue(c, "key", "value")
}
