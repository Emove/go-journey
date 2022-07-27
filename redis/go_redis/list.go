package go_redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func rpush(client *redis.Client) {
	key := "queue:test"
	client.Del(client.Context(), key)
	push := client.RPush(client.Context(), key, "001")
	result, _ := push.Result()
	fmt.Println(result)
	client.RPush(client.Context(), key, "002")
	client.RPush(client.Context(), key, "003")
	client.RPop(client.Context(), key)
	push = client.RPush(client.Context(), key, "004")
	result, _ = push.Result()
	fmt.Println(result)
}
