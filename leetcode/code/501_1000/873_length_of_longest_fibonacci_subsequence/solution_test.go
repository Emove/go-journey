package _73_length_of_longest_fibonacci_subsequence

import "testing"

func TestLenLongestFibSubseq(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			want: 5,
		},
		{
			name: "second",
			args: args{arr: []int{1, 3, 7, 11, 12, 14, 18}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LenLongestFibSubseq(tt.args.arr); got != tt.want {
				t.Errorf("LenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
