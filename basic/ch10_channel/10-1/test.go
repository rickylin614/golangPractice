package main

// channel
// Don't communicate by sharing memory; share memory by communication
// 不要通過共享內存來通信;要通過通信來共享內存

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("worker %d receiverd %d\n", id, <-c)
	}
}

// chan<- 表示只能收数据   <-chan表示只能发数据
func createWork(id int) chan<- int { //在里面才使用go建立goroutine
	c := make(chan int) //创建channel 返回出去
	go func() {
		for {
			fmt.Printf("worker %d receiverd %c\n", id, <-c)
		}
	}()
	return c
}

func channelDemo() {
	// var c chan int // c = nil
	c := make(chan int)
	// go func() { //匿名函數接收 n的寫法
	// 	for {
	// 		n := <-c
	// 		fmt.Println(n)
	// 	}
	// }()
	go worker(999, c)
	go worker(998, c) //对同个channel多定义 将会随机分配宣告的这两个
	c <- 1            //使用channel時 必須要有gorouting去接收值
	c <- 2
	c <- 3
	c <- 4

	var chans [20]chan int
	var chansRecei [20]chan<- int
	for i := 0; i < 10; i++ {
		chans[i] = make(chan int)
		go worker(i, chans[i]) //呼叫寫好的chan用func
		chansRecei[i+10] = createWork(i)
	}
	for i := 0; i < 10; i++ {
		chans[i] <- i
		chansRecei[i+10] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChanngel() {
	c := make(chan int, 3) // 3 表示缓冲 表示可以缓冲3笔数据 第四笔前可以不需要接收者（存起来）
	c <- 1
	c <- 2
	c <- 3
	go worker(500, c)
	c <- 4
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 4)
	go worker(404, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) //close前會等待所有c運算完畢
}

func main() {
	channelDemo()
	bufferedChanngel()
	channelClose()
}
