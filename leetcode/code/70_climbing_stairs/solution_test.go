package _0_climbing_stairs

import (
	"testing"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "second",
			args: args{n: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs1(tt.args.n); got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
