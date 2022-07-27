package _65_array_nesting

import "testing"

func TestArrayNesting(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{nums: []int{5, 4, 0, 3, 1, 6, 2}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayNesting(tt.args.nums); got != tt.want {
				t.Errorf("ArrayNesting() = %v, want %v", got, tt.want)
			}
		})
	}
}
