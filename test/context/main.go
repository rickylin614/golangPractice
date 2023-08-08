package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	contextFunc(&ctx)
	fmt.Println(ctx.Value("key"))
	ctx2 := context.WithValue(ctx, "key", "value2")
	fmt.Println(ctx.Value("key"))
	fmt.Println(ctx2.Value("key"))

}

func contextFunc(c *context.Context) {
	*c = context.WithValue(*c, "key", "value3")
	// fmt.Println(c)
}

func contextTest() {

}
