package _35_encode_and_decode_tinyurl

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		original := "https://leetcode.com/problems/design-tinyurl"
		codec := Constructor()
		tinyURL := codec.encode(original)
		originalURL := codec.decode(tinyURL)
		if original != originalURL {
			t.Errorf("decode fail, want: %s, got: %s", original, originalURL)
		}
	})
}
