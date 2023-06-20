package card

import (
	"testing"
)

func TestCombination(t *testing.T) {
	deck := []int{0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E}
	out := make(chan []int)
	testNumber := 5

	go func() {
		Combination([]int{}, deck, testNumber, out)
		close(out)
	}()

	expectedCombinations := choose(len(deck), testNumber)

	count := 0
	for range out {
		count++
	}

	if count != expectedCombinations {
		t.Errorf("Expected %d combinations, but got %d", expectedCombinations, count)
	}
}

func choose(n, k int) int {
	if k > n {
		return 0
	}
	if k > n/2 {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res *= n - i
		res /= i + 1
	}
	return res
}
