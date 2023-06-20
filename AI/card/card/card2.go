package card

import (
	"context"
	"sort"
	"sync"
)

type Card6Result struct {
	playerA       []Card
	playerAResult string
	playerB       []Card
	playerBReulst string
	isTie         bool
}

// 判斷6張手牌的worker,產10個goroutine並且回傳chan
func Run6CardHandelWorker(ctx context.Context, cardChan chan []int) chan Card6Result {
	resultChan := make(chan Card6Result, 1024)
	wg := sync.WaitGroup{}
	wg.Add(32)
	for i := 0; i < 32; i++ {
		go func() {
			for cards := range cardChan {
				combination := make([]int, 3)
				choose3(cards, combination, 0, 0, resultChan)
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

// 判斷三張手牌的牌型種類
func Hand3CardType(handInt []int) ([]Card, string) {
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
	for i := 0; i < len(handInt)-1; i++ {
		if hand[i].Suit != hand[i+1].Suit {
			sameSuit = false
		}
		if hand[i].Rank+1 != hand[i+1].Rank {
			rankSequence = false
		}
		sameRankCounts[hand[i].Rank]++
	}
	sameRankCounts[hand[len(handInt)-1].Rank]++

	// 特殊順子判斷 A 2 3
	if hand[0].Rank == 0x02 && hand[1].Rank == 0x03 && hand[2].Rank == 0x0E {
		rankSequence = true
	}

	// 順子同花判斷完
	switch {
	case sameSuit && rankSequence && hand[len(handInt)-1].Rank == 0x0E && hand[len(handInt)-2].Rank == 0x0D:
		return hand, "皇家同花順"
	case sameSuit && rankSequence:
		return hand, "同花順"
	case sameSuit:
		return hand, "同花"
	case rankSequence:
		return hand, "順子"
	}

	for _, count := range sameRankCounts {
		if count == 3 {
			return hand, "三條"
		}
		if count == 2 {
			return hand, "一對"
		}
	}
	return hand, "高牌"

}

func getIsTie(hand1 []Card, hand1key string, hand2 []Card, hand2key string) (isTie bool) {
	if hand1key == hand2key {
		if hand1[0].Rank == hand2[0].Rank && hand1[1].Rank == hand2[1].Rank && hand1[2].Rank == hand2[2].Rank {
			return true
		}
	}
	// if hand1key == hand2key {
	// 	if hand1key == "皇家同花順" || hand1key == "同花順" || hand1key == "順子" {
	// 		if hand1[0].Rank == hand2[0].Rank && hand1[1].Rank == hand2[1].Rank && hand1[2].Rank == hand2[2].Rank {
	// 			return true
	// 		}
	// 	} else if hand1key == "同花" || hand1key == "高牌" {

	// 		if hand1[2].Rank == hand2[2].Rank {
	// 			return true
	// 		}

	// 	} else if hand1key == "一對" {
	// 		twopairA := 0
	// 		twopairB := 0

	// 		if hand1[0].Rank == hand1[1].Rank {
	// 			twopairA = hand1[0].Rank
	// 		} else if hand1[0].Rank == hand1[2].Rank {
	// 			twopairA = hand1[0].Rank
	// 		} else if hand1[1].Rank == hand1[2].Rank {
	// 			twopairA = hand1[1].Rank
	// 		}
	// 		if hand2[0].Rank == hand2[1].Rank {
	// 			twopairB = hand2[0].Rank
	// 		} else if hand2[0].Rank == hand2[2].Rank {
	// 			twopairB = hand2[0].Rank
	// 		} else if hand2[1].Rank == hand2[2].Rank {
	// 			twopairB = hand2[1].Rank
	// 		}
	// 		if twopairA == twopairB {
	// 			return true
	// 		}
	// 	}
	// }
	return false
}

// 遞迴方法來找到所有的3張牌的組合
func choose3(cards []int, combination []int, idx, start int, resultChan chan<- Card6Result) {
	if idx == 3 {
		// 對每一個組合進行相應的處理
		playerACard, resultA := Hand3CardType(combination)
		playerBCard, resultB := Hand3CardType(getOrderCard(cards, combination))
		// 將結果傳遞回去
		resultChan <- Card6Result{
			playerA:       playerACard,
			playerAResult: resultA,
			playerB:       playerBCard,
			playerBReulst: resultB,
			isTie:         getIsTie(playerACard, resultA, playerBCard, resultB),
		}
		return
	}
	for i := start; i < len(cards); i++ {
		combination[idx] = cards[i]
		choose3(cards, combination, idx+1, i+1, resultChan)
	}
}

func getOrderCard(cards []int, combination []int) []int {
	otherCards := make([]int, 0, 3)
	for _, card := range cards {
		if !contains(combination, card) {
			otherCards = append(otherCards, card)
		}
	}
	return otherCards
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
