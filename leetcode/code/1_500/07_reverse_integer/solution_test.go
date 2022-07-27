package _7_reverse_integer

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "second",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "third",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "forth",
			args: args{x: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.x); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
