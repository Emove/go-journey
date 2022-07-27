package _5_longest_palindromic_substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "second",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "third",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "forth",
			args: args{s: "bb"},
			want: "bb",
		},
		{
			name: "fifth",
			args: args{s: "ccc"},
			want: "ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
