package _68_binary_gap

import "testing"

func TestBinaryGap(t *testing.T) {
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
			args:    args{n: 22},
			wantAns: 2,
		},
		{
			name:    "second",
			args:    args{n: 8},
			wantAns: 0,
		},
		{
			name:    "third",
			args:    args{n: 5},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := BinaryGap(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("BinaryGap() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
