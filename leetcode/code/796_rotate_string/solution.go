package _96_rotate_string

import "strings"

func RotateString(s, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
