package go_redis

import "testing"

func TestStream(t *testing.T) {
	InitGoRedis()

	XRange()
}
