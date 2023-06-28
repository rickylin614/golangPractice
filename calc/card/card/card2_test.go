package card

import (
	"testing"
)

func TestHand3CardType(t *testing.T) {
	tests := []struct {
		handInt []int
		expect  string
	}{
		{[]int{0x32, 0x13, 0x04}, "順子"},
		{[]int{0x32, 0x22, 0x13}, "一對"},
		{[]int{0x32, 0x12, 0x22}, "三條"},
		{[]int{0x31, 0x21, 0x01}, "三條"},
		{[]int{0x0A, 0x3B, 0x1C}, "順子"},
		{[]int{0x3A, 0x3B, 0x3C}, "同花順"},
		{[]int{0x3C, 0x3D, 0x3E}, "皇家同花順"},
		{[]int{0x32, 0x34, 0x36}, "同花"},
		{[]int{0x32, 0x24, 0x36}, "高牌"},
		{[]int{0x2C, 0x2D, 0x2E}, "皇家同花順"},
		{[]int{0x22, 0x1E, 0x23}, "順子"},
	}

	for _, test := range tests {
		_, result := Hand3CardType(test.handInt)
		if result != test.expect {
			t.Errorf("Hand3CardType(%v) expected %v, got %v", test.handInt, test.expect, result)
		}
	}
}

func TestGetWhoWin(t *testing.T) {
	tests := []struct {
		hand1    []Card
		hand1Key string
		hand2    []Card
		hand2Key string
		expected bool
	}{
		{
			hand1:    []Card{{0x02, 0x3}, {0x03, 0x3}, {0x04, 0x3}}, // 順子
			hand1Key: "順子",
			hand2:    []Card{{0x02, 0x2}, {0x03, 0x2}, {0x04, 0x2}}, // 順子
			hand2Key: "順子",
			expected: true,
		},
		{
			hand1:    []Card{{0x02, 0x3}, {0x03, 0x3}, {0x04, 0x3}}, // 順子
			hand1Key: "順子",
			hand2:    []Card{{0x02, 0x2}, {0x03, 0x2}, {0x05, 0x2}}, // 不是順子
			hand2Key: "高牌",
			expected: false,
		},
		{
			hand1:    []Card{{0x02, 0x0}, {0x02, 0x1}, {0x04, 0x0}}, // 一對
			hand1Key: "一對",
			hand2:    []Card{{0x02, 0x2}, {0x02, 0x3}, {0x04, 0x2}}, // 一對
			hand2Key: "一對",
			expected: true,
		},
		{
			hand1:    []Card{{0x02, 0x0}, {0x05, 0x0}, {0x09, 0x0}}, // 同花
			hand1Key: "同花",
			hand2:    []Card{{0x03, 0x2}, {0x04, 0x2}, {0x09, 0x2}}, // 同花
			hand2Key: "同花",
			expected: true,
		},
	}

	for _, test := range tests {
		result := getIsTie(test.hand1, test.hand1Key, test.hand2, test.hand2Key)
		if result != test.expected {
			t.Errorf("getwhowin(%v, %s, %v, %s) expected %t, got %t", test.hand1, test.hand1Key, test.hand2, test.hand2Key, test.expected, result)
		}
	}
}
