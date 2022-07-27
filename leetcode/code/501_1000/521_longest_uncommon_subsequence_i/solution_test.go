package _21_longest_uncommon_subsequence_i

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{a: "aba", b: "cdc"},
			want: 3,
		},
		{
			name: "second",
			args: args{a: "aaa", b: "bbb"},
			want: 3,
		},
		{
			name: "third",
			args: args{a: "aaa", b: "aaa"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLUSLength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
