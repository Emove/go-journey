package go_redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var GoRedisClient *redis.Client
var GoRedisSentinelClient *redis.SentinelClient

func InitGoRedis() {
	// 启动redis实例
	GoRedisClient = redis.NewClient(&redis.Options{
		Addr:         "192.168.1.234:6380",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     100,
		PoolTimeout:  30 * time.Second,
	})
	GoRedisSentinelClient = redis.NewSentinelClient(&redis.Options{
		Addr:         "192.168.1.234:6380",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     100,
		PoolTimeout:  30 * time.Second,
	})
	redis.NewClusterClient()
}
