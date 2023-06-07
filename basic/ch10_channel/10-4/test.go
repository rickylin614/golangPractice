package main

import (
	"fmt"
	"math/rand"
	"time"
)

//使用select 進行調度

func generator(param int) chan int {
	resp := make(chan int)
	go func() {
		i := 0
		for {
			sleepParam := time.Duration(rand.Intn(250*param)) * time.Millisecond
			time.Sleep(sleepParam)
			resp <- i
			i++
		}
	}()
	return resp
}

// func worker(id int, c chan int) {
// 	for {
// 		fmt.Printf("worker %d receiverd %d\n", id, <-c)
// 	}
// }

func createWork(id int) chan int { //在里面才使用go建立goroutine
	c := make(chan int) //创建channel 返回出去
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("worker %d receiverd %d\n", id, <-c)
		}
	}()
	return c
}

func main() {
	// var c1, c2 chan int // c1 and c2 = nil
	var c1, c2 = generator(1), generator(2)
	w := createWork(0)

	n := 0
	var values []int
	tm := time.After(15 * time.Second)
	tick := time.Tick(time.Millisecond * 50)
	// loop:
	for {
		var activeWorker chan<- int // nil channel
		var acticeValue int
		if len(values) > 0 {
			activeWorker = w
			acticeValue = values[0]
		}
		select {
		case n = <-c1: //此段的意思為 當此channel有送資料出來時 會執行此段 並使用n接收
			fmt.Println("Received from c1:", n)
			if n < 100 {
				values = append(values, n)
			}
		case n = <-c2:
			fmt.Println("Received from Ｃ２:", n)
			if n < 100 {
				values = append(values, n)
			}
		case activeWorker <- acticeValue:
			if len(values) > 0 {
				values = values[1:]
			}
		case <-tick: //固定时间印出剩馀长度
			fmt.Println("values len = ", len(values))
		case <-time.After(500 * time.Millisecond):
			fmt.Println("time out")

		case <-tm: //此段的意思為 當此channel有送資料出來時 會執行此段 tm則是會在設置好的時間 也就是15秒時將資料送出
			fmt.Println("bye")
			break
		}
	}

	// time.Sleep(time.Second)
}
