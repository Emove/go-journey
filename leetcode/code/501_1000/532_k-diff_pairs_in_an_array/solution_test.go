package _32_k_diff_pairs_in_an_array

import "testing"

func TestFindPairs(t *testing.T) {
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
			args: args{nums: []int{3, 1, 4, 1, 5}, k: 2},
			want: 2,
		},
		{
			name: "second",
			args: args{nums: []int{1, 2, 3, 4, 5}, k: 1},
			want: 4,
		},
		{
			name: "third",
			args: args{nums: []int{1, 3, 1, 5, 4}, k: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindPairs() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPairs1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
