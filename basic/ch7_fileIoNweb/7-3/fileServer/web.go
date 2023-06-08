package main

import (
	"fmt"
	"log"
	"net/http"
	"practice/basic/ch7_fileIoNweb/7-3/fileServer/myHandle"

	// _ "net/http/pprof" //網址輸入/debug/pprof可以觀察內部數據
	"os"
	// "github.com/gpmgo/gopm/log"
)

type appHandler func(writer http.ResponseWriter, req *http.Request) error

// 調用func(writer http.ResponseWriter, req *http.Request) error 並轉換為  func(http.ResponseWriter, *http.Request)
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	//headler 為一個function
	//依然需要一個匿名函數去實現 調用handler並取出func(http.ResponseWriter, *http.Request)
	return func(resp http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(resp, http.StatusText(500), 500)
			}
		}()
		var code int
		if err := handler(resp, req); err != nil {
			userErr, ok := err.(userError)
			switch {
			case ok:
				http.Error(resp, userErr.Message(), http.StatusBadRequest)
				return
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(resp, http.StatusText(code), code)
			log.Printf("Panic: %v", err)
		}
		reqMsgPrint(req)
	}
}

// 定義可以給用戶看到的ERROR 使用此類的需要有Error()以及Message兩類
type userError interface {
	error
	Message() string
}

func reqMsgPrint(req *http.Request) {
	fmt.Println("entry log")
	// for k, v := range req.Header {
	// 	values := strings.Join(v, " ")
	// 	fmt.Printf("header:%s , %s\n", k, values)
	// }
}

func main() {
	//HandleFunc需要  func(http.ResponseWriter, *http.Request) 結構
	fmt.Println("program start")
	http.HandleFunc("/", errWrapper(myHandle.HandleFileList))
	fmt.Println("setting handleFunc end")
	err := http.ListenAndServe(":8888", nil) //直接設定監聽
	fmt.Println("setting listen and serve end")
	if err != nil {
		panic(err)
	}
}
