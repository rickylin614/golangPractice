package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(NewHookSet),
		fx.Invoke(func(data *int) {
			fmt.Println("invoke func 1 ", *data)
		}),
		fx.Invoke(func(data *int) {
			fmt.Println("invoke func 2, ", *data)
		}),
		// fx.Populate(func() {
		// 	fmt.Println("populat func 1")
		// }),
		// fx.Populate(func() {
		// 	fmt.Println("populat func 2")
		// }),
	)
	// app.Start(context.Background())
	// defer app.Stop(context.Background())
	app.Run() // start and blocking

	fmt.Println("main end")
}

func NewHookSet(lc fx.Lifecycle) *int {
	i := 0
	j := 0
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			i++
			fmt.Println("Hook Onstart:", i)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			j++
			fmt.Println("Hook OnStop:", i)
			return nil
		},
	})
	return &i
}
