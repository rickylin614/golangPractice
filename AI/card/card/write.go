package card

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func WriteResultsToCSV(ctx context.Context, results <-chan CardResult, filename string) {
	// Use os.OpenFile with os.O_APPEND|os.O_CREATE|os.O_WRONLY to open the file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for result := range results {
		var row []string
		for _, card := range result.Cards {
			row = append(row, strconv.Itoa(card))
		}
		row = append(row, result.Result)
		writer.Write(row)
	}
}

func WriteCard6ResultsToCSV(ctx context.Context, results <-chan Card6Result, filename string) {
	// Use os.OpenFile with os.O_APPEND|os.O_CREATE|os.O_WRONLY to open the file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 統計使用
	playerA := make(map[string]int, 0)
	playerB := make(map[string]int, 0)
	tieCount := 0

	for result := range results {
		playerA[result.playerAResult]++
		playerB[result.playerBReulst]++
		if result.isTie {
			tieCount++
		}
	}

	fmt.Printf("玩家A牌型總計: %+v\n", playerA)
	fmt.Printf("玩家B牌型總計: %+v\n", playerB)
	fmt.Printf("和局總計: %+v\n", tieCount)
}
