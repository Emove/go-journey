package _10_minimum_height_trees

import (
	"reflect"
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{n: 4, edges: [][]int{{1, 0}, {1, 2}, {1, 3}}},
			want: []int{1},
		},
		{
			name: "second",
			args: args{n: 6, edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinHeightTrees1(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinHeightTrees2(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
