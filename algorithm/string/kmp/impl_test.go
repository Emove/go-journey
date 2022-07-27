package kmp

import "testing"

func TestKmp(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{s: "hello", p: "ll"},
			want: 2,
		},
		{
			name: "second",
			args: args{s: "hello", p: "ol"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kmp(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("Kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
