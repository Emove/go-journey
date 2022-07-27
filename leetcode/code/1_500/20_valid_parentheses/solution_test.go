package _0_valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "second",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "third",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "forth",
			args: args{s: "([)]"},
			want: false,
		},
		{
			name: "fifth",
			args: args{s: "{[]}"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
