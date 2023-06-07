package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("中文字字中文"))
	fmt.Println(lengthOfLongestSubstring("这是一句话－话里有话"))
}

func lengthOfLongestSubstring(s string) int {

	if len(s) < 2 {
		return len(s)
	}
	maxlength := 0
	startIndex := 0
	lastOccurred := make(map[rune]int) // 最後一次出現的位置
	for i, ch := range []rune(s) {
		//出現重複時 判斷重複與上一間的間格
		if _, exist := lastOccurred[ch]; exist && lastOccurred[ch] >= startIndex {
			startIndex = lastOccurred[ch] + 1
		}
		if i-startIndex+1 > maxlength {
			maxlength = i - startIndex + 1
		}
		lastOccurred[ch] = i //把此字串最後一次出現保存
	}

	return maxlength
}
