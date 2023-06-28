package card

import (
	"testing"
)

func TestHandType(t *testing.T) {
	tests := []struct {
		name     string
		hand     []int
		expected string
	}{
		{
			name:     "皇家同花順",
			hand:     []int{0x3E, 0x3D, 0x3C, 0x3B, 0x3A},
			expected: "皇家同花順",
		},
		{
			name:     "同花順",
			hand:     []int{0x32, 0x33, 0x34, 0x35, 0x36},
			expected: "同花順",
		},
		{
			name:     "順子",
			hand:     []int{0x32, 0x13, 0x24, 0x15, 0x06},
			expected: "順子",
		},
		{
			name:     "鐵支",
			hand:     []int{0x32, 0x12, 0x22, 0x33, 0x02},
			expected: "鐵支",
		},
		{
			name:     "葫蘆",
			hand:     []int{0x32, 0x22, 0x43, 0x13, 0x12},
			expected: "葫蘆",
		},
		{
			name:     "同花",
			hand:     []int{0x32, 0x34, 0x35, 0x36, 0x37},
			expected: "同花",
		},
		{
			name:     "三條",
			hand:     []int{0x32, 0x2A, 0x12, 0x19, 0x42},
			expected: "三條",
		},
		{
			name:     "兩對",
			hand:     []int{0x32, 0x22, 0x34, 0x13, 0x14},
			expected: "兩對",
		},
		{
			name:     "一對",
			hand:     []int{0x32, 0x22, 0x39, 0x14, 0x11},
			expected: "一對",
		},
		{
			name:     "高牌",
			hand:     []int{0x32, 0x24, 0x36, 0x18, 0x1A},
			expected: "高牌",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hand5CardType(tt.hand); got != tt.expected {
				t.Errorf("HandType() = %v, want %v", got, tt.expected)
			}
		})
	}
}
