package _73_matchsticks_to_square

import (
	"testing"
)

func TestMakeSquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{matchsticks: []int{1, 1, 2, 2, 2}},
			want: true,
		},
		{
			name: "second",
			args: args{matchsticks: []int{3, 3, 3, 3, 4}},
			want: false,
		},
		{
			name: "third",
			args: args{matchsticks: []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeSquare(tt.args.matchsticks); got != tt.want {
				t.Errorf("MakeSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
