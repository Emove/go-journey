package _09_fibonacci_number

import "testing"

func TestFib(t *testing.T) {
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
			want: 1,
		},
		{
			name: "second",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "third",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "forth",
			args: args{n: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
