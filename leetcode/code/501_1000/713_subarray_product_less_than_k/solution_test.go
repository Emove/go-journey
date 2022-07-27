package _13_subarray_product_less_than_k

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
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
			args: args{nums: []int{10, 5, 2, 6}, k: 100},
			want: 8,
		},
		{
			name: "second",
			args: args{nums: []int{1, 2, 3}, k: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("NumSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
