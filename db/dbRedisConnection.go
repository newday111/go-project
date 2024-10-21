package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var Rdb *redis.Client

func GetRedisConnection() {

	redisIp := viper.GetString("redis.host")
	redisPort := viper.GetString("redis.port")
	redisDb := viper.GetInt("redis.db")
	redisPool := viper.GetInt("redis.pool")

	ipPort := fmt.Sprintf("%s:%s", redisIp, redisPort)

	Rdb = redis.NewClient(&redis.Options{
		Addr:         ipPort,
		Password:     "",
		DB:           redisDb,
		PoolSize:     redisPool,
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
