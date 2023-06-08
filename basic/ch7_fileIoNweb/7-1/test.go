package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/basic/ch7_fileIoNweb/7-1/fib"
)

//資源管理及出錯處理

//defer 確保函數調用結束時發生

func tryDefer() {
	//defer先進後出 常用於close()
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	defer fmt.Println()
	for i := 0; i < 20; i++ {
		defer fmt.Printf("%d,", i)
	}
}

func WriteFile(fileName string) {
	// file, err := os.Create(fileName)
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		// panic(err)
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println("pathError:", pathError.Op, pathError.Path, pathError.Err)
		} else {
			panic(err)
		}
		// fmt.Println("file open fail:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fib()
	for i := 0; i < 20; i++ {
		// fmt.Println(f())
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	WriteFile("abc123")
}
