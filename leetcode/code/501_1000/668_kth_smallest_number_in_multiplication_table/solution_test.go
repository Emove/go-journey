package _68_kth_smallest_number_in_multiplication_table

import "testing"

func TestFindKthNumber(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{m: 3, n: 3, k: 5},
			want: 3,
		},
		{
			name: "second",
			args: args{m: 2, n: 3, k: 6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthNumber1(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("FindKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
