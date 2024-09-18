package go_redis

import "github.com/go-redis/redis/v8"

var GoRedisClient *redis.Client

func InitGoRedis() *redis.Client {
	redis.NewClient()
	ops := &redis.Options{
		Addr: "127.0.0.1:6389",
	}
	GoRedisClient = redis.NewClient(ops)
	return GoRedisClient
}
