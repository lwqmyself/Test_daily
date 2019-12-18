package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
)

/**
 * @author miku
 * @date 2019/11/30
 */
//初始化日志，使用beego框架里面的日志包
func InitLogs() (err error) {
	config := make(map[string]interface{})
	config["filename"] = "./logs/test.log"
	config["level"] = 7
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed ,err:%v", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

var RedisPool *redis.Pool

func InitRedis() (err error) {
	RedisPool = &redis.Pool{
		MaxIdle:     0,
		MaxActive:   100,
		IdleTimeout: 10,
		Dial: func() (redis.Conn, error) {
			//这里将 conn 类型赋值给接口
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

	conn := RedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("PING")
	if err != nil {
		logs.Error("ping redis failed , err:%v", err)
		return
	}
	return
}

func main() {
	InitLogs()
	InitRedis()
	conn := RedisPool.Get()
	//string
	_, err := conn.Do("set", "name", "lwq")
	if err != nil {
		logs.Error(err.Error())
	}
	res1, err := redis.String(conn.Do("get", "name"))
	fmt.Println(res1)

	//hash
	_, err = conn.Do("hmset", "user1", "age",
		10, "name", "lwq")
	if err != nil {
		logs.Error(err.Error())
	}
	res2, err := redis.Strings(conn.Do("hgetall", "user1"))
	fmt.Println(res2)

	//list
	_, err = conn.Do("del", "mylist")
	if err != nil {
		logs.Error(err.Error())
	}
	_, err = conn.Do("LPUSH", "mylist", "a", "b", "c")
	if err != nil {
		logs.Error(err.Error())
	}
	res3, err := redis.Strings(conn.Do("lrange", "mylist", 0, -1))
	fmt.Println(res3)

	//set
	_, err = conn.Do("sadd", "myset", "lwq")
	_, err = conn.Do("sadd", "myset", "hhhhhh")
	if err != nil {
		logs.Error(err.Error())
	}
	res4, err := redis.Strings(conn.Do("smembers", "myset"))
	fmt.Println(res4)

	//zset
	_, err = conn.Do("zadd", "myzset", 100, "Golang")
	_, err = conn.Do("zadd", "myzset", 66, "Java")
	if err != nil {
		logs.Error(err.Error())
	}
	//成员值递减排序
	res5, err := redis.Strings(conn.Do("ZREVRANGE", "myzset", 0, -1))
	fmt.Println(res5)
}
