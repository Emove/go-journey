package _79_largest_palindrome_product

import "testing"

func TestLargestPalindrome(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{n: 2},
			want: 987,
		},
		{
			name: "second",
			args: args{n: 1},
			want: 9,
		},
		{
			name: "third",
			args: args{n: 3},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPalindrome(tt.args.n); got != tt.want {
				t.Errorf("LargestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
