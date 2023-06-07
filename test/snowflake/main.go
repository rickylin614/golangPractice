package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	gosnowflake "github.com/godruoyi/go-snowflake"
)

func main() {

	// datas := make([]uint64, 200)
	datas := make(map[uint64]bool, 200)
	datas2 := make(map[int64]bool, 200)

	t := time.Now()
	for index := range datas {

		gosnowflake.SetStartTime(t)
		id := gosnowflake.ID()
		if _, ok := datas[index]; ok {
			datas[id] = false
		} else {
			datas[id] = true
		}
	}

	snode, _ := snowflake.NewNode(1)

	for index := range datas2 {

		id := snode.Generate().Int64()
		if _, ok := datas2[index]; ok {
			datas2[id] = false
		} else {
			datas2[id] = true
		}
	}

	fmt.Println(datas)
	fmt.Println(datas2)

}
