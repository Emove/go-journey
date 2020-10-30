package proxy

import "testing"

func TestDao_update(t *testing.T) {
	dao := NewClient()
	_, _ = dao.update()
}
