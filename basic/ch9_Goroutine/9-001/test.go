package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//goroutine 輕量級線程:協程
	//非搶占式處理，由協程主動交出控制權
	//多協程可能在多個縣程上執行

	var a [10]int
	var count int
	// a := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// count := 10
	for i := 0; i < 10; i++ {
		go func(i int) {
			for { //無限迴圈 主程序停止即停止
				// fmt.Printf("Hello from goroutine %d\n", i)
				a[i]++
				count++
				runtime.Gosched() //交出控制權 讓其餘人有機會運行 I/O select channel 等待鎖 函數調用 都是可能的切換點
				// runtime.Goexit() //直接結束程序
				// go run -race test.go 表示確認數據衝突
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	// log.Warn("Hello from goroutine %d\n", 999)
	fmt.Println(a)
	asum := 0

	for _, v := range a {
		asum += v
	}
	fmt.Println("a的總和:", asum)
	fmt.Println("程序的count++:", count)
	fmt.Println("main thread end")
}
