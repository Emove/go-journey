package _24_wiggle_sort_ii

import (
	"leetcode/code/eq"
	"testing"
)

func TestWiggleSort(t *testing.T) {
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
			args: args{nums: []int{1, 5, 1, 1, 6, 4}},
			want: []int{1, 6, 1, 5, 1, 4},
		},
		{
			name: "second",
			args: args{nums: []int{1, 2, 2, 3}},
			want: []int{2, 3, 1, 2},
		},
		{
			name: "third",
			args: args{nums: []int{1, 3, 2, 2, 3, 1}},
			want: []int{2, 3, 1, 3, 1, 2},
		},
		{
			name: "forth",
			args: args{nums: []int{1, 4, 3, 4, 1, 2, 1, 3, 1, 3, 2, 3, 3}},
			want: []int{3, 4, 2, 4, 2, 3, 1, 3, 1, 3, 1, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if WiggleSort(tt.args.nums); !eq.EqualSliceInt(tt.args.nums, tt.want) {
				t.Errorf("MinMoves2() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
