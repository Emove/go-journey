package _252_cells_with_odd_values_in_a_matrix

import "testing"

func TestOddCells(t *testing.T) {
	type args struct {
		m       int
		n       int
		indices [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{m: 2, n: 3, indices: [][]int{{0, 1}, {1, 1}}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddCells(tt.args.m, tt.args.n, tt.args.indices); got != tt.want {
				t.Errorf("OddCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
