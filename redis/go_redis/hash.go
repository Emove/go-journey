package go_redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func createKeys(client *redis.Client) {
	key1 := "topic:00001:jc_psb0001"
	client.Del(client.Context(), key1)
	client.HSet(client.Context(), key1, "key001", "value001")
	client.HSet(client.Context(), key1, "key002", "value002")

	key2 := "topic:00001:jc_psb0002"
	client.Del(client.Context(), key2)
	client.HSet(client.Context(), key2, "key003", "value003")
	client.HSet(client.Context(), key2, "key004", "value004")

	key3 := "topic:00002:jc_psb0001"
	client.Del(client.Context(), key3)
	client.HSet(client.Context(), key3, "key005", "value005")
	client.HSet(client.Context(), key3, "key006", "value006")

	key4 := "topic:00002:jc_psb0002"
	client.Del(client.Context(), key4)
	client.HSet(client.Context(), key4, "key007", "value007")
	client.HSet(client.Context(), key4, "key008", "value008")
}

func scanKeys(client *redis.Client) {
	keys := client.Keys(client.Context(), "topic:*")
	for _, key := range keys.Val() {
		scanValues(key, client)
	}
}

func scanValues(hashKey string, client *redis.Client) {
	length := client.HLen(client.Context(), hashKey)
	scan := client.HScan(client.Context(), hashKey, 0, "", length.Val())
	iterator := scan.Iterator()

	var key string
	var value string
	count := 0
	for iterator.Next(client.Context()) {
		count++
		if count%2 == 1 {
			key = iterator.Val()
		} else {
			value = iterator.Val()
			fmt.Printf("hashKey: %s, key: %s, value: %s\n", hashKey, key, value)
			client.HDel(client.Context(), hashKey, key)
			count = 0
		}
	}
}
