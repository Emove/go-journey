package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"time"
)

var RedisClient *redis.Client

func init() {
	options := &redis.Options{
		Addr: "192.168.1.2:6379",
		DB:   0,
	}
	RedisClient = redis.NewClient(options)
}

type UnLock interface {
	Unlock() bool
}

type lock struct {
	key   string
	rs    *redsync.Redsync
	mutex *redsync.Mutex
}

func Lock(key string, expire time.Duration) (UnLock, error) {
	pool := goredis.NewPool(RedisClient)
	rs := redsync.New(pool)
	l := &lock{
		key:   key,
		rs:    rs,
		mutex: rs.NewMutex(key, redsync.WithExpiry(expire)),
	}

	return l, l.mutex.Lock()
}

func (l *lock) Unlock() bool {
	if _, err := l.mutex.Unlock(); err != nil {
		return false
	}
	return true
}
