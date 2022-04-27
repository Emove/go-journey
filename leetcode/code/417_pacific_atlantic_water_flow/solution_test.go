package _17_pacific_atlantic_water_flow

import (
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{heights: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}},
			want: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name: "second",
			args: args{heights: [][]int{{2, 1}, {1, 2}}},
			want: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name: "third",
			args: args{heights: [][]int{{3, 3, 3}, {3, 1, 3}, {0, 2, 4}}},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PacificAtlantic(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PacificAtlantic() = %v, want %v", got, tt.want)
			}
		})
	}
}
