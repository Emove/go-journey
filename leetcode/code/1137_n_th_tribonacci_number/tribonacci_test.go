package n_th_tribonacci_number_1137

import "testing"

func TestTribonacci(t *testing.T) {
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
			args: args{n: 3},
			want: 2,
		},
		{
			name: "second",
			args: args{n: 4},
			want: 4,
		},
		{
			name: "third",
			args: args{n: 25},
			want: 1389537,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tribonacci(tt.args.n); got != tt.want {
				t.Errorf("Tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
