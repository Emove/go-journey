package _1_two_sum

import (
	"leetcode/code/eq"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{nums: []int{2, 7, 11, 15}, target: 9},
			want: []int{0, 1},
		},
		{
			name: "second",
			args: args{nums: []int{3, 2, 4}, target: 6},
			want: []int{1, 2},
		},
		{
			name: "third",
			args: args{nums: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !eq.EqualSliceInt(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
