package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	contextFunc(&ctx)
	fmt.Println(ctx.Value("key"))
	ctx2 := context.WithValue(ctx, "key", "value2")
	fmt.Println(ctx.Value("key"))
	fmt.Println(ctx2.Value("key"))

	ctx3, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-ctx3.Done()
		fmt.Println("ctx3 done")
	}()
	// 確認ctx4的cancel是否會觸發ctx3
	ctx4, cancel := context.WithTimeout(ctx3, time.Second*3)

	go func() {
		cancel()
	}()

	select {
	case <-ctx4.Done():
		fmt.Println("ctx4 done")
		// case <-time.Tick(time.Second * 4):
		// 	fmt.Println("time tick")
	}
	time.Sleep(time.Second * 10)

}

func contextFunc(c *context.Context) {
	*c = context.WithValue(*c, "key", "value3")
	// fmt.Println(c)
}

func contextTest() {

}
