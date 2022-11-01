package _2_longest_valid_parentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "first",
			args:    args{s: "(()"},
			wantAns: 2,
		},
		{
			name:    "second",
			args:    args{s: ")()())"},
			wantAns: 4,
		},
		{
			name:    "third",
			args:    args{s: ""},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := LongestValidParentheses(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("LongestValidParentheses() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
