package _64_can_i_win

import "testing"

func TestCanIWin(t *testing.T) {
	type args struct {
		maxChoosableInteger int
		desiredTotal        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{maxChoosableInteger: 10, desiredTotal: 11},
			want: false,
		},
		{
			name: "second",
			args: args{maxChoosableInteger: 10, desiredTotal: 0},
			want: true,
		},
		{
			name: "third",
			args: args{maxChoosableInteger: 10, desiredTotal: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanIWin(tt.args.maxChoosableInteger, tt.args.desiredTotal); got != tt.want {
				t.Errorf("CanIWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
