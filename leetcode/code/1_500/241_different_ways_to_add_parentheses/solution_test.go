package _41_different_ways_to_add_parentheses

import (
	"leetcode/code/eq"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{expression: "2-1-1"},
			want: []int{0, 2},
		},
		{
			name: "second",
			args: args{expression: "2*3-4*5"},
			want: []int{-34, -14, -10, -10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffWaysToCompute(tt.args.expression); !eq.EqualSliceIntElement(got, tt.want) {
				t.Errorf("DiffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
