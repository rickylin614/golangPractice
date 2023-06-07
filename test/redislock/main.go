package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

var client *redis.Client
var clusterClient *redis.ClusterClient
var isCluster bool = false
var rs *redsync.Redsync

func main() {
	NewRedis("127.0.0.1:6379")

	ml := rs.NewMutex("ricky")

	if err := ml.Lock(); err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		b, err := ml.Unlock()
		fmt.Println(b, err)
	}()
	time.Sleep(time.Second * 20)
	fmt.Println(time.Now().String())
}

func NewRedis(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	// 確認連線正常
	if _, err := rdb.Ping(context.TODO()).Result(); err != nil {
		log.Panic(err)
	}

	// 創建redsync
	pool := goredis.NewPool(rdb)
	rs = redsync.New(pool)
	isCluster = false

	return rdb
}
