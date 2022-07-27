package _29_consecutive_numbers_sum

import "testing"

func TestConsecutiveNumbersSum(t *testing.T) {
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
			args:    args{n: 5},
			wantAns: 2,
		},
		{
			name:    "second",
			args:    args{n: 9},
			wantAns: 3,
		},
		{
			name:    "third",
			args:    args{n: 15},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := ConsecutiveNumbersSum(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("ConsecutiveNumbersSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
