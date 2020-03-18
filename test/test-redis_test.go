package test

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T) {

	//创建一个redis连接
	c, err := redis.Dial("tcp", "10.9.40.193:6379", redis.DialPassword("Clt@123456"))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	//defer作用：本函数执行完后执行c.Close()
	defer c.Close()

	//设置一个缓存数据
	_, err = c.Do("SET", "go_key", "hello world")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	//检查是否存在缓存数据
	is_key_exit, err := redis.Bool(c.Do("EXISTS", "go_key"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}

	//获取一个缓存数据
	username, err := redis.String(c.Do("GET", "go_key"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get go_key: %v \n", username)
	}

	//删除一个缓存数据
	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	} else {
		fmt.Println("delete sucessed")
	}

	// redis list操作
	_, err = c.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

}
