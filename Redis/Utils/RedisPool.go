package Utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"

	"time"
)

var (
	RedisPool *redis.Pool
)

func InitRedis() (err error) {
	RedisPool = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   0,
		IdleTimeout: time.Duration(10) * time.Second,
		Dial: func() (redis.Conn, error) {

			return redis.Dial("tcp", "localhost:6379")
		},
	}

	conn := RedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("PING")
	if err != nil {
		logs.Error("ping redis failed , err:%v", err)
		return
	}
	fmt.Println("redis 初始化成功")
	return
}
