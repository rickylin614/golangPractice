package main

import (
	"card/card"
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	Run6Card()
	fmt.Println(time.Since(t))
}

// 只計算每種牌型產生總量
func Run6Card() {
	ctx := context.Background()
	gen := card.RunGenCardWorker(ctx, 6) // 預計產六張手牌所有組合
	resultChan := card.Run6CardHandelWorker(ctx, gen)
	card.WriteCard6ResultsToCSV(ctx, resultChan, "test.csv")
}

// 列出每種手牌及牌型並寫入檔案中
func Run5Card() {
	ctx := context.Background()
	gen := card.RunGenCardWorker(ctx, 5) // 預計產五張手牌所有組合
	resultChan := card.Run5CardHandelWorker(ctx, gen)
	card.WriteResultsToCSV(ctx, resultChan, "test.csv")
}
