package main

import (
	"Test_daily/Redis/Utils"
	"fmt"
	"go-common/library/cache/redis"
)

func main() {
	Utils.InitRedis()
	pool := Utils.RedisPool
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "lwq")
	if err != nil {
		fmt.Println(err.Error())
	}
	res1, err := redis.String(conn.Do("get", "name"))
	fmt.Println(res1)
}
