package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Init() *redis.Client {
	// 创建一个 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 密码，如果没有设置密码则为空
		DB:       0,                // 使用默认的 DB
	})

	// 创建一个上下文
	ctx := context.Background()

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("connect Redis fail:", err)
		return nil
	}
	fmt.Println("connect Redis successful:", pong)

	return rdb
}
