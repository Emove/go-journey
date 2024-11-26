package string

import (
	"strings"
)

func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

func sliceString(s string) string {
	return s[1:]
}
