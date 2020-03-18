package test

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func TestRedisClient(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "10.9.40.193:6379",
		Password: "Clt@123456", // no password set
		DB:       0,            // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// 设置一个缓存
	err = client.Set("go-client-key", "hello redis client", 0).Err()
	if err != nil {
		panic(err)
	}

	//得到一个缓存
	val, err := client.Get("go-client-key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("go-client-key：", val)

	//得到一个缓存
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
