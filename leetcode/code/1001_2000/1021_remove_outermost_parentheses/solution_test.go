package _021_remove_outermost_parentheses

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
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
			args: args{s: "(()())(())"},
			want: "()()()",
		},
		{
			name: "second",
			args: args{s: "(()())(())(()(()))"},
			want: "()()()()(())",
		},
		{
			name: "third",
			args: args{s: "()()"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveOuterParentheses(tt.args.s); got != tt.want {
				t.Errorf("RemoveOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
