package go_redis

import "fmt"

func XRange() {
	xRange := GoRedisClient.XRange(GoRedisClient.Context(), "test_stream", "-", "+")
	val := xRange.Val()
	for _, msg := range val {
		fmt.Printf(msg.ID + "----")
		for key, value := range msg.Values {
			fmt.Printf(key+"=%s\n", value)
		}
	}
}
