package _58_find_k_closest_elements

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "second",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
