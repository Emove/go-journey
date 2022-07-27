package _8_implement_strstr

import "testing"

func TestStrStr(t *testing.T) {
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
			args: args{s: "aaaaa", p: "bba"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
