package move_zeroes_283

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{nums: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			for i, val := range tt.args.nums {
				if val != tt.want[i] {
					t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
				}
			}
		})
	}
}
