package model

import (
	"context"
	"fmt"
	"time"
	"vvblog/config"
	"vvblog/vlog"

	"github.com/go-redis/redis/v8"
)

var (
	RDB *redis.Client
)

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Secret,
		DB:       int(config.Redis.Database),
		PoolSize: 10,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		vlog.Fatalf("连接Redis出错：%v", err)
	}
}
