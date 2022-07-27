package _98_diagonal_traverse

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "second",
			args: args{mat: [][]int{{1, 2}, {3, 4}}},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDiagonalOrder(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
