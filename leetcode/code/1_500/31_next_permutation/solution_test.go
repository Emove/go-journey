package _1_next_permutation

import (
	"leetcode/code/eq"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 3, 2},
		},
		{
			name: "second",
			args: args{nums: []int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "third",
			args: args{nums: []int{1, 1, 5}},
			want: []int{1, 5, 1},
		},
		{
			name: "forth",
			args: args{nums: []int{1, 2}},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NextPermutation(tt.args.nums)
			if !eq.EqualSliceInt(tt.args.nums, tt.want) {
				t.Errorf("want=%v, got=%v", tt.want, tt.args.nums)
			}
		})
	}
}
