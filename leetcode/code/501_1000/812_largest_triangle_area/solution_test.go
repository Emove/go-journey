package _12_largest_triangle_area

import (
	"testing"
)

func TestLargestTriangleArea(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "first",
			args: args{points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}},
			want: float64(2),
		},
		{
			name: "second",
			args: args{points: [][]int{{4, 6}, {6, 5}, {3, 1}}},
			want: 5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestTriangleArea(tt.args.points); got != tt.want {
				t.Errorf("LargestTriangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
