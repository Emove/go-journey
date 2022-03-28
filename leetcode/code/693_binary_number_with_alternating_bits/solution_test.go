package _93_binary_number_with_alternating_bits

import (
	"testing"
)

func TestHasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{n: 5},
			want: true,
		},
		{
			name: "second",
			args: args{n: 7},
			want: false,
		},
		{
			name: "third",
			args: args{n: 11},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAlternatingBits(tt.args.n); got != tt.want {
				t.Errorf("HasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
