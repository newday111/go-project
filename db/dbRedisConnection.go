package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rdb *redis.Client

func GetRedisConnection() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.0:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})
	ctx := context.Background()
	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功", pong)
}
