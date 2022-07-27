package _99_falling_squares

import (
	"reflect"
	"testing"
)

func TestFallingSquares(t *testing.T) {
	type args struct {
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{positions: [][]int{{1, 2}, {2, 3}, {6, 1}}},
			want: []int{2, 5, 5},
		},
		{
			name: "second",
			args: args{positions: [][]int{{100, 100}, {200, 100}}},
			want: []int{100, 100},
		},
		{
			name: "third",
			args: args{positions: [][]int{{4, 9}, {8, 8}, {6, 8}, {8, 2}, {1, 2}}},
			want: []int{9, 17, 25, 27, 27},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FallingSquares1(tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FallingSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
