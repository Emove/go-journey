package _96_rotate_string

import "testing"

func TestRotateString(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{s: "abcde", goal: "cdeab"},
			want: true,
		},
		{
			name: "second",
			args: args{s: "abcde", goal: "abced"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("RotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
