package _175_prime_arrangements

import "testing"

func TestNumPrimeArrangements(t *testing.T) {
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
			args: args{n: 5},
			want: 12,
		},
		{
			name: "second",
			args: args{n: 100},
			want: 682289015,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumPrimeArrangements(tt.args.n); got != tt.want {
				t.Errorf("NumPrimeArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
