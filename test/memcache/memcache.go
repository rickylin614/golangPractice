package memcache

// go tool compile -S .\memcache.go > output.txt
// 結論 宣告變數當記憶體空間放沒有浪費執行效率 編譯的時候自己會優化效能 詳細: output.txt

//go:noinline
func Cal(i int) int {
	m := map[int]int{
		1: 2,
		4: 6,
		9: 10,
	}

	if v, ok := m[i]; ok {
		return v
	}
	return i
}

func m1() int {
	result := 0
	for i := 0; i < 100000; i++ {
		result += Cal(i) * 2
	}
	return result
}

var c int

func m2() int {
	result := 0
	for i := 0; i < 100000; i++ {
		c = Cal(i)
		result += c * 2
	}
	return result
}

func m3() int {
	result := 0
	for i := 0; i < 100000; i++ {
		result += i*2 + 6
	}
	return result
}
