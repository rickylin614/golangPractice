package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	// data := "HappyNewYear"
	// _ = ioutil.WriteFile(filename, []byte(data), os.ModeAppend)
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// if err != nil { //直接條件式寫在宣告
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
	)
}

func grade(score int) string { //go語言函數返回值寫在後面
	g := ""
	switch {
	case score > 100 || score < 0:
		panic(fmt.Sprintf("Wroung socre %d", score)) //panic程式會直接中斷 並抱位置
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score >= 90:
		g = "A"
	}
	return g
}

// func bounded() {

// }
