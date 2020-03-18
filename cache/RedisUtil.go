package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "10.9.40.193:6379",
		Password: "Clt@123456", // no password set
		DB:       0,            // use default DB
	})
	pong, err := Client.Ping().Result()
	fmt.Println("redis client init", pong, err)
}

func GetOne(redisKey string) string {
	//得到一个缓存
	val, err := Client.Get(redisKey).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func SetOne(redisKey string, value string) {
	//得到一个缓存
	_, err := Client.Set(redisKey, value, 0).Result()
	if err != nil {
		panic(err)
	}
}
