package _089_duplicate_zeros

import (
	"leetcode/code/eq"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{arr: []int{1, 0, 2, 3, 0, 4, 5, 0}},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "second",
			args: args{[]int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DuplicateZeros(tt.args.arr)
			if !eq.EqualSliceInt(tt.args.arr, tt.want) {
				t.Errorf("DeplicateZeros()=%v, wamt=%v", tt.args.arr, tt.want)
			} else {
				t.Logf("DeplicateZeros()=%v, wamt=%v", tt.args.arr, tt.want)
			}
		})
	}
}
