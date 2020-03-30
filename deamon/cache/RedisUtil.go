package cache

/**
created by hbd
2020-03-18
Redis工具类
*/

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

/**
初始化Redis连接
*/
func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "10.9.40.193:6379",
		Password: "Clt@123456", // no password set
		DB:       0,            // use default DB
	})
	pong, err := RedisClient.Ping().Result()
	fmt.Println("redis client init", pong, err)
}

func GetOne(redisKey string) string {
	//得到一个缓存
	val, err := RedisClient.Get(redisKey).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func SetOne(redisKey string, value string) {
	//得到一个缓存
	_, err := RedisClient.Set(redisKey, value, 0).Result()
	if err != nil {
		panic(err)
	}
}

/**
发布
*/
func Publish(channel string, message string) int64 {
	val, err := RedisClient.Publish(channel, message).Result()
	if err != nil {
		panic(err)
	}
	return val
}

/**
订阅
*/
func Subscribe(channel string) {

	sub := RedisClient.Subscribe(channel)

	//开启一个线程侦听channel
	go func() {
		var receipt interface{}
		var err error
		for {
			//receive 阻塞式调用
			receipt, err = sub.Receive()
			if err != nil {
				fmt.Println(err)
			}
			if receipt != "" {
				switch v := receipt.(type) {
				case *redis.Message: //单个订阅subscribe
					fmt.Printf("%s: message: %s\n", v.Channel, v.Payload)
				case error:
					return
				}
			}
		}
	}()
}

/**
map 新增元素
*/
func HSet(redisKey string, hashKey string, hashValue string) int64 {
	val, err := RedisClient.HSet(redisKey, hashKey, hashValue).Result()
	if err != nil {
		panic(err)
	}
	return val
}

/**
map 移除一个元素
*/
func HDel(redisKey string, hashKey string) int64 {
	val, err := RedisClient.HDel(redisKey, hashKey).Result()
	if err != nil {
		panic(err)
	}
	return val
}

/**
map 获取元素
*/
func HGet(redisKey string, hashKey string) string {
	b, err := RedisClient.HExists(redisKey, hashKey).Result()
	if err != nil {
		panic(err)
	}
	if b {
		val, err := RedisClient.HGet(redisKey, hashKey).Result()
		if err != nil {
			panic(err)
		}
		return val
	}
	return ""
}

/**
向列表头部插入元素
*/
func LPush(redisKey string, value string) {
	err := RedisClient.LPush(redisKey, value).Err()
	if err != nil {
		panic(err)
	}
}

/**
列表读取某一个位置的元素
*/
func LIndex(redisKey string, index int64) string {
	val, err := RedisClient.LIndex(redisKey, index).Result()
	if err != nil {
		panic(err)
	}
	return val
}

/**
列表删除元素
*/
func LRemove(redisKey string, value string) int64 {
	val, err := RedisClient.LRem(redisKey, 0, value).Result()
	if err != nil {
		panic(err)
	}
	return val
}

/**
列表拿取元素
*/
func LPop(redisKey string) string {
	val, err := RedisClient.LPop(redisKey).Result()
	if err != nil {
		panic(err)
	}
	return val
}
