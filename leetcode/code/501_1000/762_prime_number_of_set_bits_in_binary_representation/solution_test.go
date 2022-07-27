package _62_prime_number_of_set_bits_in_binary_representation

import (
	"testing"
)

func TestCountPrimeSetBits(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "first",
			args:    args{left: 6, right: 10},
			wantAns: 4,
		},
		{
			name:    "second",
			args:    args{left: 10, right: 15},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := CountPrimeSetBits(tt.args.left, tt.args.right); gotAns != tt.wantAns {
				t.Errorf("CountPrimeSetBits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
