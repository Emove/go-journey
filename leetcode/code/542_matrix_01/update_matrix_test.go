package matrix_01_542

import (
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			name: "second",
			args: args{matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UpdateMatrix(tt.args.matrix)
			want := tt.want
			for i := 0; i < len(want); i++ {
				for j := 0; j < len(want[0]); j++ {
					if want[i][j] != got[i][j] {
						t.Errorf("%s test failed, value not match, want: %d, got: %d, index: %d %d", tt.name, want[i][j], got[i][j], i, j)
					}
				}
			}
		})
	}
}
