package Utils

import (
	"fmt"
	"go-common/library/cache/redis"
)

func nitRedis() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis 连接失败 err : ", err.Error())
		return
	}
	defer conn.Close()
	//conn.Do("ping")
	_, err = conn.Do("SET", "name", "lwq")
	if err != nil {
		fmt.Println("set 失败 err : ", err.Error())
		return
	}
	//错误写法 nameString := r.(string)
	r, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("get name err : ", err.Error())
		return
	}

	fmt.Println(r)

	_, err = conn.Do("hmset", "user1", "age",
		10, "name", "lwq")
	if err != nil {
		fmt.Println(err.Error())
	}
	res1, err := redis.Strings(conn.Do("hgetall", "user1"))

	fmt.Println(res1)
	fmt.Println("redis 连接成功")

}
