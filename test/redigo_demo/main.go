package main

import (
	"fmt"
	"time"

	"log"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}

func Increx(pool *redis.Pool, key interface{}, timeoutInSeconds int) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	lua := redis.NewScript(1, `
		local num = redis.call('INCR',KEYS[1])
		if num == 1 then
			redis.call('EXPIRE',KEYS[1],ARGV[1])
		end
		return num
	`)

	incrementedValue, err := redis.Int(lua.Do(conn, key, timeoutInSeconds))
	if err != nil {
		log.Fatal("redis", fmt.Sprintf("pool : %v, err : %v, key : %v", pool, err, key))
		return 0, err
	}

	return incrementedValue, nil
}

func main() {
	a, _ := Increx(pool, "a", 10)
	b, _ := Increx(pool, "a", 10)
	c, _ := Increx(pool, "a", 10)
	fmt.Println(a, b, c)
	a, _ = Increx(pool, "a", 10)
	time.Sleep(time.Second * 11)
	b, _ = Increx(pool, "a", 10)
	c, _ = Increx(pool, "a", 10)
	fmt.Println(a, b, c)
}
