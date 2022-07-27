package _72_factorial_trailing_zeroes

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "first",
			args:    args{n: 3},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := trailingZeroes(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("trailingZeroes() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
