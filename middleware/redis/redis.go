package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var rdb *redis.Client

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database string
}

func InitRedis(reConfig *RedisConfig) *redis.Client {
	fmt.Println("==============init redis start==============")
	rdb = redis.NewClient(&redis.Options{
		Addr:     reConfig.Host + ":" + strconv.Itoa(reConfig.Port),
		Password: reConfig.Password,
		DB:       0,
	})
	ping, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(ping)
	}
	fmt.Println("==============init redis success==============")
	return rdb
}

func GetRedis() *redis.Client {
	return rdb
}

func SetKey(key string, value interface{}, expiration time.Duration) {
	res, err := rdb.Set(context.Background(), key, value, expiration).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}

func GetKey(key string) (string, error) {
	res, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func DelKey(key string) {
	res := rdb.Del(context.Background(), key)
	if res.Err() != nil {
		fmt.Println(res.Err())
	}
}
