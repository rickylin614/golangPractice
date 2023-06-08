package main

import (
	"fmt"
	"practice/basic/ch5_interface/5-1/mock"
	"practice/basic/ch5_interface/5-1/real"
	"time"
)

// interface 內部表達的即函數
// 實現為隱式 實現者只需定義結構(可不同)以及實現的方法
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// func post(poster Poster) {
// 	poster.Post("http://www.imooc.com", map[string]string{
// 		"name":   "ccmouse",
// 		"course": "golang",
// 	})
// }

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another imooc.com"})
	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake imooc.com"}
	fmt.Println(download(r))
	inspect(r)
	fmt.Printf("%T %v\n", r, r)
	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	// fmt.Println(download(r))
	inspect(r)
	fmt.Printf("%T %v\n", r, r)

	if realRetriever, ok := r.(real.Retriever); ok { // ()括號內表示類型 此種寫法表示轉態
		fmt.Println(realRetriever.UserAgent)
	}

	rr := mock.Retriever{Contents: "this is a fake imooc.com"}
	fmt.Println(session(&rr)) //此處呼叫的變數 裡面需要擁有get post兩種方法
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:" + v.Contents)
	case real.Retriever:
		fmt.Println("UserAgent:" + v.UserAgent)
		fmt.Printf("TimeOut:%v\n", v.TimeOut.Seconds())
	}
}
