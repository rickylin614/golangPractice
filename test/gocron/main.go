package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	// 設定所有job的最多同時工作數量
	//  - RescheduleMode: 超過數量則略過此次
	//  - WaitMode: 超過數量等待此次結束
	s.SetMaxConcurrentJobs(10, gocron.WaitMode)

	i := 0
	s.Every(1).Seconds().Do(HelloWorld, &i)
	s.Every(1).Seconds().Do(HelloHeaven)
	// SingletonMode: 此job同時間最多只會執行一次
	s.Every(1).SingletonMode().Seconds().Do(SleepFiveSecounds)
	s.Every(1).Day().At("10:30").At("08:00").Do(HelloDays)
	s.StartAsync() // work啟動

	// 监听操作系统的中断信号（Ctrl+C）
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt, os.Kill)

	// 阻塞等待中断信号
	<-osSignal
	s.Stop()
	fmt.Println("End")
}

func HelloWorld(times *int) {
	*times++
	fmt.Println("HelloWorld", *times)
}

func HelloHeaven() {
	fmt.Println("HelloHeaven")
}

func SleepFiveSecounds() {
	time.Sleep(time.Second * 5)
	fmt.Println("Five Second")
}

func HelloDays() {
	fmt.Println("HelloDays")
}
