package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var rdb redis.UniversalClient
var key = "ChannelKey"

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		// Addr: "localhost:49153",
	})
	resp := rdb.Ping(context.TODO())
	if resp.Err() != nil {
		panic(resp.Err())
	}
	go SubRedis()
}

func main() {
	InitRedis()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/push", func(ctx *gin.Context) {
		r := rdb.Publish(ctx, key, "this is a message")
		if r.Err() != nil {
			fmt.Println(r.Err())
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "push!",
		})
	})

	r.Run()

}

func SubRedis() {
	c := rdb.Subscribe(context.TODO(), key)
	for {
		msg := <-c.Channel()
		fmt.Println(msg)
	}
}
