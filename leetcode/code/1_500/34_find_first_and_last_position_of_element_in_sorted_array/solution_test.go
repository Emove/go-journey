package _4_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
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
			args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 8},
			want: []int{3, 4},
		},
		{
			name: "second",
			args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 6},
			want: []int{-1, -1},
		},
		{
			name: "third",
			args: args{nums: []int{}, target: 0},
			want: []int{-1, -1},
		},
		{
			name: "forth",
			args: args{nums: []int{2, 2}, target: 2},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
