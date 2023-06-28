package card

import (
	"context"
	"sort"
	"sync"
)

type CardResult struct {
	Cards  []int  // 手牌
	Result string // 牌型
}

// 產10個goroutine並且回傳chan
func Run5CardHandelWorker(ctx context.Context, cardChan chan []int) chan CardResult {
	resultChan := make(chan CardResult, 1024)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for cards := range cardChan {
				r := CardResult{
					Cards:  cards,
					Result: Hand5CardType(cards),
				}
				resultChan <- r
			}
			wg.Done()
		}()
	}

	// 確認做完關閉resultChan
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

type Card struct {
	Rank int
	Suit int
}

func NewCard(value int) Card {
	return Card{
		Suit: value & 0xf0,
		Rank: value & 0x0f,
	}
}

// 判斷五張手牌的牌型種類
func Hand5CardType(handInt []int) string {
	hand := make([]Card, len(handInt))
	for i, cardInt := range handInt {
		hand[i] = NewCard(cardInt)
	}

	sort.Slice(hand, func(i, j int) bool {
		return hand[i].Rank < hand[j].Rank
	})

	var sameSuit = true
	var sameRankCounts = make(map[int]int)
	var rankSequence = true
	for i := 0; i < 4; i++ {
		if hand[i].Suit != hand[i+1].Suit {
			sameSuit = false
		}
		if hand[i].Rank+1 != hand[i+1].Rank {
			rankSequence = false
		}
		sameRankCounts[hand[i].Rank]++
	}
	sameRankCounts[hand[4].Rank]++

	// 順子同花判斷完
	switch {
	case sameSuit && rankSequence && hand[len(handInt)-1].Rank == 0x0E:
		return "皇家同花順"
	case sameSuit && rankSequence:
		return "同花順"
	case sameSuit:
		return "同花"
	case rankSequence:
		return "順子"
	}

	var fourOfAKind = false
	var threeOfAKind = make(map[int]bool)
	var pair = make(map[int]bool)
	for card, count := range sameRankCounts {
		if count >= 4 {
			fourOfAKind = true
			break
		}
		if count == 3 {
			threeOfAKind[card] = true
		}
		if count == 2 {
			pair[card] = true
		}
	}

	switch {
	case fourOfAKind:
		return "鐵支"
	case len(threeOfAKind) == 1 && len(pair) == 1:
		return "葫蘆"
	case len(threeOfAKind) == 1:
		return "三條"
	case len(pair) >= 2:
		return "兩對"
	case len(pair) == 1:
		return "一對"
	default:
		return "高牌"
	}
}
