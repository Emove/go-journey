package _051_height_checker

import "testing"

func TestHeightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{heights: []int{1, 1, 4, 2, 1, 3}},
			want: 3,
		},
		{
			name: "second",
			args: args{heights: []int{5, 1, 2, 3, 4}},
			want: 5,
		},
		{
			name: "third",
			args: args{heights: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeightChecker(tt.args.heights); got != tt.want {
				t.Errorf("HeightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
