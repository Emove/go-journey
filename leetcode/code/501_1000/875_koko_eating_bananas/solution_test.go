package _75_koko_eating_bananas

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{piles: []int{3, 6, 7, 11}, h: 8},
			want: 4,
		},
		{
			name: "second",
			args: args{piles: []int{30, 11, 23, 4, 20}, h: 5},
			want: 30,
		},
		{
			name: "third",
			args: args{piles: []int{30, 11, 23, 4, 20}, h: 6},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("MinEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
