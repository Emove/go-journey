package art

import (
	"testing"
)

func Test_find(t *testing.T) {
	set("redis", "redis")
	set("elastic", "elastis")

	if i, b := find("elastic"); b {
		t.Log(i)
	}
}
