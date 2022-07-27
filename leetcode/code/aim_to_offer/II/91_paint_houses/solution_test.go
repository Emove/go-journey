package _1_paint_houses

import "testing"

func TestMinCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{costs: [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCost(tt.args.costs); got != tt.want {
				t.Errorf("MinCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
