package _5_three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "second",
			args: args{nums: []int{-2, 0, 0, 2, 2}},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "third",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "forth",
			args: args{nums: []int{0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
