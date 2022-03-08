package _055_plates_between_candles

import (
	"reflect"
	"testing"
)

func TestPlatesBetweenCandles(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{s: "**|**|***|", queries: [][]int{{2, 5}, {5, 9}}},
			want: []int{2, 3},
		},
		{
			name: "second",
			args: args{s: "***|**|*****|**||**|*", queries: [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}},
			want: []int{9, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlatesBetweenCandles(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlatesBetweenCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
