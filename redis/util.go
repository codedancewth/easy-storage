package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// redis的通用工具

type ReClient struct {
	log *log.Logger
	ctx context.Context
	rc  *redis.Client
}

func NewReClient() *ReClient {
	rc := Init()
	return &ReClient{
		ctx: context.Background(),
		rc:  rc,
	}
}

// Set
// timeDuration is seconds
func (r *ReClient) Set(key string, value interface{}, timeDuration int64) (bool, error) {
	result, err := r.rc.Set(r.ctx, key, value, time.Duration(timeDuration)*time.Second).Result()
	if err != nil {
		log.Fatalln(fmt.Sprintf("set err %v", err))
		return false, err
	}
	fmt.Println(result)
	if result == "OK" {
		return true, err
	}
	return false, err
}

// Get 获取用户的值，支持反序列化
func (r *ReClient) Get(key string, value interface{}) (interface{}, error) {
	result, err := r.rc.Get(r.ctx, key).Result()
	if err != nil {
		log.Fatalln(fmt.Sprintf("set err %v", err))
		return value, err
	}
	fmt.Println(result)
	// 反序列化
	if err = json.Unmarshal([]byte(result), &value); err != nil {
		log.Fatalln(fmt.Sprintf("get value Unmarshal err %v", err))
		return value, err
	}
	return value, err
}
