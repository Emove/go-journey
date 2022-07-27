package _19_find_k_th_smallest_pair_distance

import "testing"

func TestSmallestDistancePair(t *testing.T) {
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
			args: args{nums: []int{1, 3, 1}, k: 1},
			want: 0,
		},
		{
			name: "second",
			args: args{nums: []int{1, 1, 1}, k: 2},
			want: 0,
		},
		{
			name: "third",
			args: args{nums: []int{1, 6, 1}, k: 3},
			want: 5,
		},
		{
			name: "forth",
			args: args{nums: []int{38, 33, 57, 65, 13, 2, 86, 75, 4, 56}, k: 26},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestDistancePair(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("SmallestDistancePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
