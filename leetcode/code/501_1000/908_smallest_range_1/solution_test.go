package _08_smallest_range_1

import "testing"

func TestSmallestRangeI(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{nums: []int{1}, k: 0},
			want: 0,
		},
		{
			name: "second",
			args: args{nums: []int{0, 10}, k: 2},
			want: 6,
		},
		{
			name: "third",
			args: args{nums: []int{1, 3, 6}, k: 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestRangeI(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("SmallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}
