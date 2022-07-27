package _20_triangle

import (
	"fmt"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
			want: 11,
		},
		{
			name: "second",
			args: args{triangle: [][]int{{-10}}},
			want: -10,
		},
		{
			name: "third",
			args: args{triangle: [][]int{{-1}, {-2, -3}}},
			want: -4,
		},
		{
			name: "forth",
			args: args{triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}},
			want: -1,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := MinimumTotal(tt.args.triangle); got != tt.want {
	//			t.Errorf("MinimumTotal() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := DPMinimumTotal(tt.args.triangle); got != tt.want {
	//			t.Errorf("MinimumTotal() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptimizeDp(tt.args.triangle); got != tt.want {
				t.Errorf("MinimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	fmt.Println(128 & 192)
	fmt.Println(min(-2, -3))
}
