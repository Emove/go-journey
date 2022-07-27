package _200_minimum_absolute_difference

import (
	"reflect"
	"testing"
)

func TestMinimumAbsDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{arr: []int{4, 2, 1, 3}},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name: "second",
			args: args{arr: []int{1, 3, 6, 10, 15}},
			want: [][]int{{1, 3}},
		},
		{
			name: "third",
			args: args{arr: []int{3, 8, -10, 23, 19, -4, -14, 27}},
			want: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumAbsDifference(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimumAbsDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
