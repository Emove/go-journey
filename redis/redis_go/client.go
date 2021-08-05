package redis_go

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisGoClient redis.Conn

func InitRedisGoClient() {
	var f = func(db int) *redis.Pool {
		return &redis.Pool{
			MaxIdle:     1,
			MaxActive:   1,
			IdleTimeout: time.Duration(5) * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%s", "192.168.1.234:6380"))
				if err != nil {
					panic(err)
				}

				_, err = c.Do("SELECT", db)
				if err != nil {
					panic(err)
				}
				return c, nil
			},
		}
	}
	RedisGoClient = f(0).Get()
}
