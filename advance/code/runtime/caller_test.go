package runtime

import "testing"

func TestGetCaller(t *testing.T) {
	t.Logf("caller: %s", GetCaller(0))
}
