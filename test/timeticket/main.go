package main

import (
	"fmt"
	"time"
)

func main() {

	dbchan := make(chan bool)

	go func() {
		// db query
		time.Sleep(6 * time.Second)
		dbchan <- true
	}()

	t := time.After(time.Second * 5)

	for {
		select {
		case <-dbchan:
			// do nothing
		case <-t:
			fmt.Println("123")
			// log print
			<-dbchan // 若要繼續等待，則打此行，若不等待，則此行替代成回傳錯誤。
		}

	}

}
