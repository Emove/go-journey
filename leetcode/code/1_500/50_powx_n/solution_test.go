package _0_powx_n

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "first",
			args: args{x: 2.0, n: 10},
			want: 1024.0,
		},
		{
			name: "second",
			args: args{x: 2.1, n: 3},
			want: 9.261,
		},
		{
			name: "third",
			args: args{x: 2.0, n: -2},
			want: 0.25000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
