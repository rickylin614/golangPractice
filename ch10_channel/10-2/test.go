package main

// channel
// Don't communicate by sharing memory; share memory by communication
// 不要通過共享內存來通信;要通過通信來共享內存

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.c {
		// n <- c
		fmt.Printf("worker %d receiverd %d\n", id, n)
		w.done()
	}
}

type worker struct {
	c    chan int
	done func()
}

func createWork(id int, wg *sync.WaitGroup) worker {
	c := make(chan int, 2)
	work := worker{c: c, done: func() {
		wg.Done()
	}}
	go doWorker(id, work)
	return work
}

func channelDemo() {
	var wg sync.WaitGroup
	wg.Add(20)

	var w [10]worker
	for i := 0; i < 10; i++ {
		w[i] = createWork(i, &wg)
	}

	for i := 0; i < 10; i++ {
		w[i].c <- i
	}
	for i := 0; i < 10; i++ {
		w[i].c <- (i + 10)
	}
	// //wait all end
	// for _, workOne := range w {
	// 	<-workOne.done //程序就會等待做好 才下一次
	// 	<-workOne.done
	// }
	wg.Wait()
}

func main() {
	channelDemo()
}
