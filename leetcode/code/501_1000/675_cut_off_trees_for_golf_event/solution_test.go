package _75_cut_off_trees_for_golf_event

import (
	"testing"
)

func TestCutOffTree(t *testing.T) {
	type args struct {
		forest [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{forest: [][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}},
			want: 6,
		},
		{
			name: "second",
			args: args{forest: [][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}}},
			want: -1,
		},
		{
			name: "third",
			args: args{forest: [][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutOffTree(tt.args.forest); got != tt.want {
				t.Errorf("CutOffTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
