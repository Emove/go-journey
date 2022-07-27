package _97_rotate_function

import "testing"

func TestMaxRotateFunction(t *testing.T) {
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
			args: args{nums: []int{4, 3, 2, 6}},
			want: 26,
		},
		{
			name: "second",
			args: args{nums: []int{100}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxRotateFunction(tt.args.nums); got != tt.want {
				t.Errorf("MaxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
