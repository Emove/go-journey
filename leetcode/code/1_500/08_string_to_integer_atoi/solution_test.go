package _8_string_to_integer_atoi

import "testing"

func TestMyAtoi(t *testing.T) {
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
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "second",
			args: args{s: "   -42"},
			want: -42,
		},
		{
			name: "third",
			args: args{s: "4193 with words"},
			want: 4193,
		},
		{
			name: "forth",
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			name: "fifth",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "sixth",
			args: args{s: "+-12"},
			want: 0,
		},
		{
			name: "seventh",
			args: args{s: "-+12"},
			want: 0,
		},
		{
			name: "eighth",
			args: args{s: "  0000000000012345678"},
			want: 12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi(tt.args.s); got != tt.want {
				t.Errorf("MyAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
