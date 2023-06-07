package main

import (
	"testing"
)

//visual studio code go插件 自動產生的測試單元
func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args //寫入的參數
		want int  //希望正確的回傳值
	}{
		//這一區塊為自定義區塊
		{name: "1", args: args{s: "aaabbb"}, want: 2},
		{name: "2", args: args{s: "bbbbb"}, want: 1},
		{name: "3", args: args{s: "pwwkew"}, want: 3},
		{name: "4", args: args{s: ""}, want: 0},
		{name: "5", args: args{s: "中文字字中文"}, want: 3},
		{name: "6", args: args{s: "这是一句话－话里有话"}, want: 6},
		{"7", args{"gabcdefggg"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("name:= %s lengthOfLongestSubstring() = %v, data = %s, want %v", tt.name, got, tt.args.s, tt.want)
			}
		})
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "这是一句话－话里有话"
	want := 6
	actual := lengthOfLongestSubstring(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if want != actual {
			b.Errorf("lengthOfLongestSubstring() = %v, data = %s, want %v", actual, s, want)
		}
	}

}
