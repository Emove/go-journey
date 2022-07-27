package _3_roman_to_integer

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "second",
			args: args{s: "IV"},
			want: 4,
		},
		{
			name: "third",
			args: args{s: "IX"},
			want: 9,
		},
		{
			name: "forth",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "fifth",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInt(tt.args.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
