package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	str := t.In(time.UTC).Format("2006-01-02 15:04:05")
	fmt.Println(str)
	loc, err := time.LoadLocation("Asia/Shanghai")
	fmt.Println(loc, err)
	if err == nil {
		fmt.Println(t.In(loc).Format("2006-01-02 15:04:05"))
	}

	loc = time.FixedZone("UTC+8", 8*3600)
	fmt.Println(t.In(loc).Format("2006-01-02 15:04:05"))
	// loc2 :=

}
