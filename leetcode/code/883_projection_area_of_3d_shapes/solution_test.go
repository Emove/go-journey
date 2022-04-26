package _83_projection_area_of_3d_shapes

import "testing"

func TestProjectionArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{grid: [][]int{{1, 2}, {3, 4}}},
			want: 17,
		},
		{
			name: "second",
			args: args{grid: [][]int{{2}}},
			want: 5,
		},
		{
			name: "third",
			args: args{grid: [][]int{{1, 0}, {0, 2}}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProjectionArea(tt.args.grid); got != tt.want {
				t.Errorf("ProjectionArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
